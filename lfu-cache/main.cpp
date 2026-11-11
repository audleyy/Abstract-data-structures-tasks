#include <iostream>
#include <string>

#include "core/lfu.h"

int main() {
	int capacity;
	int queryCount;
	if (!(cin >> capacity >> queryCount) || capacity < 0 || queryCount < 0) {
		return 1;
	}

	LFUCache cache(capacity);
	bool firstResult = true;

	for (int i = 0; i < queryCount; ++i) {
		string command;
		int key;

		if (!(cin >> command >> key)) return 1;

		if (command == "SET") {
			int value;
			if (!(cin >> value)) return 1;
			cache.set(key, value);
		} else if (command == "GET") {
			if (!firstResult) cout << ' ';
			cout << cache.get(key);
			firstResult = false;
		} else {
			return 1;
		}
	}

	cout << '\n';
	return 0;
}
