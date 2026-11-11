#ifndef LFU_CACHE_H
#define LFU_CACHE_H

using namespace std;

struct CacheEntry {
	int key;
	int value;
	int frequency;
	int lastUsed;
};

class LFUCache {
private:
	CacheEntry* entries;
	int capacity;
	int size;
	int time;

	int findKey(int key) const {
		for (int i = 0; i < size; ++i) {
			if (entries[i].key == key) return i;
		}
		return -1;
	}

	int findLeastUsed() const {
		int leastUsed = 0;
		for (int i = 1; i < size; ++i) {
			if (entries[i].frequency < entries[leastUsed].frequency ||
				(entries[i].frequency == entries[leastUsed].frequency &&
				 entries[i].lastUsed < entries[leastUsed].lastUsed)) {
				leastUsed = i;
			}
		}
		return leastUsed;
	}

	void recordAccess(int index) {
		++entries[index].frequency;
		entries[index].lastUsed = ++time;
	}

public:
	LFUCache(int capacity) {
		if (capacity < 0) capacity = 0;
		this->capacity = capacity;
		size = 0;
		time = 0;
		entries = new CacheEntry[capacity]{};
	}

	void set(int key, int value) {
		if (capacity == 0) return;

		int index = findKey(key);
		if (index != -1) {
			entries[index].value = value;
			recordAccess(index);
			return;
		}

		if (size < capacity) {
			index = size;
			++size;
		} else {
			index = findLeastUsed();
		}

		entries[index] = CacheEntry{key, value, 1, ++time};
	}

	int get(int key) {
		int index = findKey(key);
		if (index == -1) return -1;

		recordAccess(index);
		return entries[index].value;
	}

	~LFUCache() {
		delete[] entries;
	}
};

#endif
