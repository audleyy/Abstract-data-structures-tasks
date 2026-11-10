#ifndef SET_H
#define SET_H

#include <iostream>
#include <string>

using namespace std;

struct Node {
	string value;
	Node* next;
};

struct Set {
	Node** buckets;
	size_t capacity;
	int size;
};

void init(Set* set);
size_t hashFunction(const string &value, size_t capacity);
bool contains(const Set* set, const string &val);
void rehash(Set* set);
void add(Set* set, const string& value);
bool remove(Set* set, const string& val);
void clear(Set* set);

#endif