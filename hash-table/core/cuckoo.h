#ifndef CUCKOO_H
#define CUCKOO_H

#include <string>
#include <utility>
#include "bloom.h"

using namespace std;

struct CuckooCell {
	std::pair<string, string> data;
	bool occupied;
};

struct Cuckoo {
	BloomFilter filter;
	CuckooCell* first;
	CuckooCell* second;
	int capacity;
	int size;
	int seed;
};

void initCuckoo(Cuckoo* table);
bool containsCuckoo(const Cuckoo* table, const string& key);
const string* getCuckoo(const Cuckoo* table, const string& key);
void addCuckoo(Cuckoo* table, const string& key);
void putCuckoo(Cuckoo* table, const string& key, const string& value);
bool removeCuckoo(Cuckoo* table, const string& key);
void clearCuckoo(Cuckoo* table);

#endif
