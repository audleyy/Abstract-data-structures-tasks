#include "hash.h"

size_t hashFunction(const string& value, size_t capacity) {
	size_t hash = 0;
	for (unsigned char symbol : value) {
		hash = hash * 31 + symbol;
	}
	return hash % capacity;
}

int firstCuckooHash(const string& value, int capacity, int seed) {
	const size_t prime = 1000000007;
	size_t multiplier = 31 + seed;
	size_t hash = 0;

	for (unsigned char symbol : value) {
		hash = (hash * multiplier + symbol) % prime;
	}
	return hash % capacity;
}

int secondCuckooHash(const string& value, int capacity, int seed) {
	const size_t prime = 1000000007;
	size_t multiplier = 32 + seed;
	size_t hash = 0;

	for (unsigned char symbol : value) {
		hash = (hash * multiplier + symbol) % prime;
	}
	return hash % capacity;
}
