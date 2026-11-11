#include "bloom.h"
#include <stdexcept>

const int hashCount = 3;

void calculateBloomHashes(const string& key, size_t& first, size_t& second) {
	first = 0;
	second = 0;

	for (unsigned char symbol : key) {
		first = first * 31 + symbol;
		second = second * 37 + symbol;
	}

	first ^= first >> 16;
	second ^= second >> 16;
	second = second * 2 + 1;
}

void initBloomFilter(BloomFilter* filter, int bitCount) {
	if (bitCount <= 0) return;
	initFilledArray(filter->bits, bitCount, "0");
}

void addToBloomFilter(BloomFilter* filter, const string& key) {
	size_t first, second;
	calculateBloomHashes(key, first, second);

	for (int i = 0; i < hashCount; ++i) {
		int index = (first + i * second) % filter->bits.size;
		replace(filter->bits, index, "1");
	}
}

bool mightContainBloomFilter(const BloomFilter* filter, const string& key) {
	if (filter->bits.size == 0) return false;

	size_t first, second;
	calculateBloomHashes(key, first, second);

	for (int i = 0; i < hashCount; ++i) {
		int index = (first + i * second) % filter->bits.size;
		if (get(filter->bits, index) == "0") return false;
	}
	
	return true;
}

void clearBloomFilter(BloomFilter* filter) {
	destroy(filter->bits);
}
