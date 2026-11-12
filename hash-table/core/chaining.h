#ifndef CHAINING_H
#define CHAINING_H

#include <cstddef>
#include <string>
#include <utility>
#include "bloom.h"

using namespace std;

struct Node {
	std::pair<string, string> data;
	Node* next;
};

struct Chaining {
	BloomFilter filter;
	Node** buckets;
	size_t capacity;
	int size;
};

void initChaining(Chaining* table);
bool containsChaining(const Chaining* table, const string& key);
const string* getChaining(const Chaining* table, const string& key);
void rehashChaining(Chaining* table);
void addChaining(Chaining* table, const string& key);
void putChaining(Chaining* table, const string& key, const string& value);
bool removeChaining(Chaining* table, const string& key);
void clearChaining(Chaining* table);

#endif
