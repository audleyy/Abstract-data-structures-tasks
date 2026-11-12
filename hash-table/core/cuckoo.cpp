#include "cuckoo.h"
#include "../../hash-functions/hash.h"
#include <utility>

void initCuckoo(Cuckoo* table, int capacity, int seed) {
	table->first = new CuckooCell[capacity]{};
	table->second = new CuckooCell[capacity]{};
	table->capacity = capacity;
	table->size = 0;
	table->seed = seed;
	initBloomFilter(&table->filter, capacity * 16);
}

void initCuckoo(Cuckoo* table) {
	initCuckoo(table, 2, 1);
}

CuckooCell* findCell(const Cuckoo* table, const string& key) {
	if (!mightContainBloomFilter(&table->filter, key)) return nullptr;

	CuckooCell& first = table->first[firstCuckooHash(key, table->capacity, table->seed)];
	if (first.occupied && first.data.first == key) return &first;

	CuckooCell& second = table->second[secondCuckooHash(key, table->capacity, table->seed)];
	if (second.occupied && second.data.first == key) return &second;

	return nullptr;
}

bool containsCuckoo(const Cuckoo* table, const string& key) {
	return findCell(table, key) != nullptr;
}

const string* getCuckoo(const Cuckoo* table, const string& key) {
	const CuckooCell* cell = findCell(table, key);
	return cell == nullptr ? nullptr : &cell->data.second;
}

bool placeEntry(Cuckoo* table, CuckooCell& entry) {
	bool first = true;
	for (int step = 0; step < 32; ++step) {
		CuckooCell* cells = table->first;
		int index = firstCuckooHash(entry.data.first, table->capacity, table->seed);

		if (!first) {
			cells = table->second;
			index = secondCuckooHash(entry.data.first, table->capacity, table->seed);
		}

		CuckooCell& cell = cells[index];
		swap(cell, entry);
		if (!entry.occupied) {
			return true;
		}
		first = !first;
	}

	return false;
}

bool copyEntries(const Cuckoo* source, Cuckoo* target) {
	for (int i = 0; i < source->capacity; ++i) {
		if (source->first[i].occupied) {
			CuckooCell entry = source->first[i];
			if (!placeEntry(target, entry)) return false;
		}
		if (source->second[i].occupied) {
			CuckooCell entry = source->second[i];
			if (!placeEntry(target, entry)) return false;
		}
	}
	return true;
}

void rebuild(Cuckoo* table, int capacity, const CuckooCell& extra) {
	int seed = table->seed + 2;
	int attempts = 0;

	while (true) {
		Cuckoo next;
		initCuckoo(&next, capacity, seed);
		CuckooCell entry = extra;
		bool placed = copyEntries(table, &next) && placeEntry(&next, entry);

		if (placed) {
			next.size = table->size + 1;
			for (int i = 0; i < next.capacity; ++i) {
				if (next.first[i].occupied) addToBloomFilter(&next.filter, next.first[i].data.first);
				if (next.second[i].occupied) addToBloomFilter(&next.filter, next.second[i].data.first);
			}
			clearCuckoo(table);
			*table = next;
			return;
		}

		clearCuckoo(&next);
		seed += 2;
		++attempts;
		if (attempts == 4) {
			capacity *= 2;
			attempts = 0;
		}
	}
}

void insertCuckoo(Cuckoo* table, const string& key, const string& value) {
	CuckooCell pending{{key, value}, true};
	if (table->size + 1 > table->capacity * 4 / 5) {
		rebuild(table, table->capacity * 2, pending);
		return;
	}

	addToBloomFilter(&table->filter, key);
	if (placeEntry(table, pending)) {
		++table->size;
	} else {
		rebuild(table, table->capacity, pending);
	}
}

void addCuckoo(Cuckoo* table, const string& key) {
	if (!containsCuckoo(table, key)) insertCuckoo(table, key, "");
}

void putCuckoo(Cuckoo* table, const string& key, const string& value) {
	CuckooCell* cell = findCell(table, key);
	if (cell != nullptr) {
		cell->data.second = value;
		return;
	}
	insertCuckoo(table, key, value);
}

bool removeCuckoo(Cuckoo* table, const string& key) {
	CuckooCell* cell = findCell(table, key);
	if (cell == nullptr) return false;

	*cell = CuckooCell{};
	--table->size;
	return true;
}

void clearCuckoo(Cuckoo* table) {
	clearBloomFilter(&table->filter);
	delete[] table->first;
	delete[] table->second;
	*table = Cuckoo{};
}
