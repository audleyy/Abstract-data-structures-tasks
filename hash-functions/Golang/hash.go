package hashfunctions

func HashFunction(value string, capacity int) int {
	var hash uint
	for i := 0; i < len(value); i++ {
		hash = hash*31 + uint(value[i])
	}
	return int(hash % uint(capacity))
}

func CuckooIndex(value string, capacity int, seed int) int {
	const prime = 1000000007
	multiplier := uint(31 + seed)
	var hash uint

	for i := 0; i < len(value); i++ {
		hash = (hash*multiplier + uint(value[i])) % prime
	}
	return int(hash % uint(capacity))
}
