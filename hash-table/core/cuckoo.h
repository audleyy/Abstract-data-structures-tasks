#ifndef CUCKOO_H
#define CUCKOO_H

#include <string>
#include "bloom.h"

using namespace std;

struct CuckooCell {
	string key;
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
void addCuckoo(Cuckoo* table, const string& key);
bool removeCuckoo(Cuckoo* table, const string& key);
void clearCuckoo(Cuckoo* table);

#endif
