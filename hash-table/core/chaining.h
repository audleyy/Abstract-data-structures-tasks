#ifndef CHAINING_H
#define CHAINING_H

#include <cstddef>
#include <string>
#include "bloom.h"

using namespace std;

struct Node {
	string value;
	Node* next;
};

struct Chaining {
	BloomFilter filter;
	Node** buckets;
	size_t capacity;
	int size;
};

void initChaining(Chaining* table);
bool containsChaining(const Chaining* table, const string& value);
void rehashChaining(Chaining* table);
void addChaining(Chaining* table, const string& value);
bool removeChaining(Chaining* table, const string& value);
void clearChaining(Chaining* table);

#endif
