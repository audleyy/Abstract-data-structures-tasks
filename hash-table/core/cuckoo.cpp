#include "cuckoo.h"
#include "../../hash-functions/hash.h"

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

bool containsCuckoo(const Cuckoo* table, const string& key) {
	if (!mightContainBloomFilter(&table->filter, key)) return false;
	const CuckooCell& first = table->first[firstCuckooHash(key, table->capacity, table->seed)];
	if (first.occupied && first.key == key) return true;

	const CuckooCell& second = table->second[secondCuckooHash(key, table->capacity, table->seed)];
	return second.occupied && second.key == key;
}

bool placeKey(Cuckoo* table, string& key) {
	bool first = true;
	for (int step = 0; step < 32; ++step) {
		CuckooCell* cells = table->first;
		int index = firstCuckooHash(key, table->capacity, table->seed);
		if (!first) {
			cells = table->second;
			index = secondCuckooHash(key, table->capacity, table->seed);
		}
		CuckooCell& cell = cells[index];

		if (!cell.occupied) {
			cell.key = key;
			cell.occupied = true;
			return true;
		}

		string displaced = cell.key;
		cell.key = key;
		key = displaced;
		first = !first;
	}

	return false;
}

bool copyKeys(const Cuckoo* source, Cuckoo* target) {
	for (int i = 0; i < source->capacity; ++i) {
		if (source->first[i].occupied) {
			string key = source->first[i].key;
			if (!placeKey(target, key)) return false;
		}
		if (source->second[i].occupied) {
			string key = source->second[i].key;
			if (!placeKey(target, key)) return false;
		}
	}
	return true;
}

void rebuild(Cuckoo* table, int capacity, const string& extra) {
	int seed = table->seed + 2;
	int attempts = 0;

	while (true) {
		Cuckoo next;
		initCuckoo(&next, capacity, seed);
		string key = extra;
		bool placed = copyKeys(table, &next) && placeKey(&next, key);

		if (placed) {
			next.size = table->size + 1;
			for (int i = 0; i < next.capacity; ++i) {
				if (next.first[i].occupied) addToBloomFilter(&next.filter, next.first[i].key);
				if (next.second[i].occupied) addToBloomFilter(&next.filter, next.second[i].key);
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

void addCuckoo(Cuckoo* table, const string& key) {
	if (containsCuckoo(table, key)) return;
	if (table->size + 1 > table->capacity * 4 / 5) {
		rebuild(table, table->capacity * 2, key);
		return;
	}

	string pending = key;
	if (placeKey(table, pending)) {
		++table->size;
		addToBloomFilter(&table->filter, key);
	} else {
		rebuild(table, table->capacity, pending);
	}
}

bool removeCuckoo(Cuckoo* table, const string& key) {
	if (!mightContainBloomFilter(&table->filter, key)) return false;
	CuckooCell* cell = &table->first[firstCuckooHash(key, table->capacity, table->seed)];
	if (!cell->occupied || cell->key != key) {
		cell = &table->second[secondCuckooHash(key, table->capacity, table->seed)];
		if (!cell->occupied || cell->key != key) return false;
	}

	cell->key = "";
	cell->occupied = false;
	--table->size;
	return true;
}

void clearCuckoo(Cuckoo* table) {
	clearBloomFilter(&table->filter);
	delete[] table->first;
	delete[] table->second;
	*table = Cuckoo{};
}
