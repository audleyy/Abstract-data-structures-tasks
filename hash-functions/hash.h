#ifndef HASH_H
#define HASH_H

#include <cstddef>
#include <string>

using namespace std;

size_t hashFunction(const string& value, size_t capacity);
int firstCuckooHash(const string& value, int capacity, int seed);
int secondCuckooHash(const string& value, int capacity, int seed);

#endif
