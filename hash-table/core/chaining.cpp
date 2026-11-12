#include "chaining.h"
#include "../../hash-functions/hash.h"

void initChaining(Chaining* table) {
	table->capacity = 4;
	table->size = 0;
	table->buckets = new Node*[table->capacity]{};
	initBloomFilter(&table->filter, table->capacity * 16);
}

Node* findNode(const Chaining* table, const string& key) {
	if (!mightContainBloomFilter(&table->filter, key)) return nullptr;
	size_t index = hashFunction(key, table->capacity);
	Node* current = table->buckets[index];

	while (current != nullptr) {
		if (current->data.first == key) return current;
		current = current->next;
	}

	return nullptr;
}

bool containsChaining(const Chaining* table, const string& key) {
	return findNode(table, key) != nullptr;
}

const string* getChaining(const Chaining* table, const string& key) {
	const Node* node = findNode(table, key);
	return node == nullptr ? nullptr : &node->data.second;
}

void rehashChaining(Chaining* table) {
	size_t newCapacity = table->capacity * 2;
	Node** newBuckets = new Node*[newCapacity]{};
	clearBloomFilter(&table->filter);
	initBloomFilter(&table->filter, newCapacity * 16);

	for (size_t i = 0; i < table->capacity; ++i) {
		Node* current = table->buckets[i];

		while (current != nullptr) {
			Node* next = current->next;

			size_t newIndex = hashFunction(current->data.first, newCapacity);
			addToBloomFilter(&table->filter, current->data.first);

			current->next = newBuckets[newIndex];
			newBuckets[newIndex] = current;

			current = next;
		}
	}

	delete[] table->buckets;
	table->buckets = newBuckets;
	table->capacity = newCapacity;
}

void insertChaining(Chaining* table, const string& key, const string& value) {
	if (table->size + 1 > table->capacity * 3 / 4) rehashChaining(table);

	size_t index = hashFunction(key, table->capacity);
	table->buckets[index] = new Node{{key, value}, table->buckets[index]};
	table->size++;
	addToBloomFilter(&table->filter, key);
}

void addChaining(Chaining* table, const string& key) {
	if (!containsChaining(table, key)) insertChaining(table, key, "");
}

void putChaining(Chaining* table, const string& key, const string& value) {
	Node* node = findNode(table, key);
	if (node != nullptr) {
		node->data.second = value;
		return;
	}
	insertChaining(table, key, value);
}

bool removeChaining(Chaining* table, const string& key) {
	if (!mightContainBloomFilter(&table->filter, key)) return false;
	size_t index = hashFunction(key, table->capacity);

	Node* current = table->buckets[index];
	Node* previous = nullptr;

	while (current != nullptr) {
		if (current->data.first == key) {
			Node* toDelete = current;

			if (previous == nullptr) {
				table->buckets[index] = toDelete->next;
			} else {
				previous->next = toDelete->next;
			}

			delete toDelete;
			table->size--;
			return true;
		}
		previous = current;
		current = current->next;
	}

	return false;
}

void clearChaining(Chaining* table) {
	clearBloomFilter(&table->filter);
	for (size_t i = 0; i < table->capacity; ++i) {
		Node* current = table->buckets[i];

		while (current != nullptr) {
			Node* next = current->next;
			delete current;
			current = next;
		}
	}

	delete[] table->buckets;
	table->capacity = 0;
	table->size = 0;
	table->buckets = nullptr;
}
