package hashtable

const hashCount = 3

type BloomFilter struct {
	bits []bool
}

func calculateBloomHashes(key string) (uint, uint) {
	var first, second uint

	for i := 0; i < len(key); i++ {
		symbol := uint(key[i])
		first = first*31 + symbol
		second = second*37 + symbol
	}

	first ^= first >> 16
	second ^= second >> 16
	second = second*2 + 1
	return first, second
}

func initBloomFilter(filter *BloomFilter, bitCount int) {
	if bitCount <= 0 {
		return
	}
	filter.bits = make([]bool, bitCount)
}

func addToBloomFilter(filter *BloomFilter, key string) {
	first, second := calculateBloomHashes(key)

	for i := uint(0); i < hashCount; i++ {
		index := (first + i*second) % uint(len(filter.bits))
		filter.bits[index] = true
	}
}

func MightContainBloomFilter(filter *BloomFilter, key string) bool {
	if len(filter.bits) == 0 {
		return false
	}

	first, second := calculateBloomHashes(key)

	for i := uint(0); i < hashCount; i++ {
		index := (first + i*second) % uint(len(filter.bits))
		if !filter.bits[index] {
			return false
		}
	}
	return true
}

func clearBloomFilter(filter *BloomFilter) {
	filter.bits = nil
}
