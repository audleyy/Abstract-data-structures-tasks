#ifndef SET_H
#define SET_H

#include "../../hash-table/core/chaining.h"

struct Set {
	Chaining table;
};

void init(Set* set);
bool contains(const Set* set, const string& value);
void rehash(Set* set);
void add(Set* set, const string& value);
bool remove(Set* set, const string& value);
void clear(Set* set);

#endif
