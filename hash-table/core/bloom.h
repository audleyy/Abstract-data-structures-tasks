#ifndef BLOOM_H
#define BLOOM_H

#include "../../array-based/core/array.h"

struct BloomFilter {
	Array bits;
};

void initBloomFilter(BloomFilter* filter, int bitCount);
void addToBloomFilter(BloomFilter* filter, const string& key);
bool mightContainBloomFilter(const BloomFilter* filter, const string& key);
void clearBloomFilter(BloomFilter* filter);

#endif
