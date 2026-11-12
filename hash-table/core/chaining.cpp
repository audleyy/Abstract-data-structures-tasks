#include "chaining.h"
#include "../../hash-functions/hash.h"

void initChaining(Chaining* table) {
	table->capacity = 4;
	table->size = 0;
	table->buckets = new Node*[table->capacity]{};
	initBloomFilter(&table->filter, table->capacity * 16);
}

bool containsChaining(const Chaining* table, const string& val) {
	if (!mightContainBloomFilter(&table->filter, val)) return false;
	size_t index = hashFunction(val, table->capacity);
	Node* current = table->buckets[index];

	while (current != nullptr) {
		if (current->value == val) return true;
		current = current->next;
	}

	return false;
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

			size_t newIndex = hashFunction(current->value, newCapacity);
			addToBloomFilter(&table->filter, current->value);

			current->next = newBuckets[newIndex];
			newBuckets[newIndex] = current;

			current = next;
		}
	}

	delete[] table->buckets;
	table->buckets = newBuckets;
	table->capacity = newCapacity;
}

void addChaining(Chaining* table, const string& value) {
	if (containsChaining(table, value)) return;

	if (table->size + 1 > table->capacity * 3 / 4) rehashChaining(table);

	size_t index = hashFunction(value, table->capacity);
	Node* newNode = new Node{value, nullptr};

	newNode->next = table->buckets[index];
	table->buckets[index] = newNode;
	table->size++;
	addToBloomFilter(&table->filter, value);
}

bool removeChaining(Chaining* table, const string& val) {
	if (!mightContainBloomFilter(&table->filter, val)) return false;
	size_t index = hashFunction(val, table->capacity);

	Node* current = table->buckets[index];
	Node* previous = nullptr;

	while (current != nullptr) {
		if (current->value == val) {
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
