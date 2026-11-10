#include "set.h"

void init(Set* set) {
	set->capacity = 4;
	set->size = 0;
	set->buckets = new Node*[set->capacity]{};
}

size_t hashFunction(const string &value, size_t capacity) {
	size_t hash = 0;

	for (unsigned char symbol : value) {
		hash = hash * 31 + symbol;
	}

	return hash % capacity;
}

bool contains(const Set* set, const string &val) {
	size_t index = hashFunction(val, set->capacity);
	Node* current = set->buckets[index];

	while (current != nullptr) {
		if (current->value == val) return true;
		current = current->next;
	}

	return false;
}

void rehash(Set* set) {
	size_t newCapacity = set->capacity * 2;
	Node** newBuckets = new Node*[newCapacity]{};

	for (size_t i = 0; i < set->capacity; ++i) {
		Node* current = set->buckets[i];

		while (current != nullptr) {
			Node* next = current->next;

    	size_t newIndex = hashFunction(current->value, newCapacity);

			current->next = newBuckets[newIndex];
			newBuckets[newIndex] = current;

			current = next;
		}
	}

	delete[] set->buckets;
	set->buckets = newBuckets;
	set->capacity = newCapacity;
}

void add(Set* set, const string& value) {
	if (contains(set, value)) return;

	if (set->size + 1 > set->capacity * 3 / 4) rehash(set);

	size_t index = hashFunction(value, set->capacity);
	Node* newNode = new Node{value, nullptr};

	newNode->next = set->buckets[index];
	set->buckets[index] = newNode;
	set->size++;
}

bool remove(Set* set, const string& val) {
	size_t index = hashFunction(val, set->capacity);

	Node* current = set->buckets[index];
	Node* previous = nullptr;

	while (current != nullptr) {
		if (current->value == val) {
			Node* toDelete = current;

			if (previous == nullptr) {
				set->buckets[index] = toDelete->next;
			} else {
				previous->next = toDelete->next;
			}

			delete toDelete;
			set->size--;
			return true;
		}
		previous = current;
		current = current->next;
	}

	return false;
}

void clear(Set* set) {
	for (size_t i = 0; i < set->capacity; ++i) {
		Node* current = set->buckets[i];
		
		while (current != nullptr) {
			Node* next = current->next;
			delete current;
			current = next;
		}
	}

	delete[] set->buckets;
	set->capacity = 0;
	set->size = 0;
	set->buckets = nullptr;
}