#include "set.h"

void init(Set* set) {
	initChaining(&set->table);
}

bool contains(const Set* set, const string& value) {
	return containsChaining(&set->table, value);
}

void rehash(Set* set) {
	rehashChaining(&set->table);
}

void add(Set* set, const string& value) {
	addChaining(&set->table, value);
}

bool remove(Set* set, const string& value) {
	return removeChaining(&set->table, value);
}

void clear(Set* set) {
	clearChaining(&set->table);
}
