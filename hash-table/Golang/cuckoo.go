package hashtable

import "lab2/hash-functions/Golang"

type CuckooCell struct {
	key      string
	occupied bool
}

type Cuckoo struct {
	Filter   BloomFilter
	first    []CuckooCell
	second   []CuckooCell
	Capacity int
	Size     int
	seed     int
}

func initCuckooCapacity(table *Cuckoo, capacity int, seed int) {
	*table = Cuckoo{}
	table.first = make([]CuckooCell, capacity)
	table.second = make([]CuckooCell, capacity)
	table.Capacity = capacity
	table.seed = seed
	initBloomFilter(&table.Filter, capacity*16)
}

func InitCuckoo(table *Cuckoo) {
	initCuckooCapacity(table, 2, 1)
}

func ContainsCuckoo(table *Cuckoo, key string) bool {
	if !MightContainBloomFilter(&table.Filter, key) {
		return false
	}
	first := table.first[hashfunctions.CuckooIndex(key, table.Capacity, table.seed)]
	if first.occupied && first.key == key {
		return true
	}

	second := table.second[hashfunctions.CuckooIndex(key, table.Capacity, table.seed+1)]
	return second.occupied && second.key == key
}

func placeKey(table *Cuckoo, key *string) bool {
	first := true
	for step := 0; step < 32; step++ {
		cells := table.first
		seed := table.seed
		if !first {
			cells = table.second
			seed++
		}
		cell := &cells[hashfunctions.CuckooIndex(*key, table.Capacity, seed)]
		if !cell.occupied {
			cell.key = *key
			cell.occupied = true
			return true
		}
		cell.key, *key = *key, cell.key
		first = !first
	}

	return false
}

func copyKeys(source *Cuckoo, target *Cuckoo) bool {
	for i := 0; i < source.Capacity; i++ {
		if source.first[i].occupied {
			key := source.first[i].key
			if !placeKey(target, &key) {
				return false
			}
		}
		if source.second[i].occupied {
			key := source.second[i].key
			if !placeKey(target, &key) {
				return false
			}
		}
	}
	return true
}

func rebuild(table *Cuckoo, capacity int, extra string) {
	seed := table.seed + 2
	attempts := 0

	for {
		var next Cuckoo
		initCuckooCapacity(&next, capacity, seed)
		key := extra
		placed := copyKeys(table, &next) && placeKey(&next, &key)
		if placed {
			next.Size = table.Size + 1
			for i := 0; i < next.Capacity; i++ {
				if next.first[i].occupied {
					addToBloomFilter(&next.Filter, next.first[i].key)
				}
				if next.second[i].occupied {
					addToBloomFilter(&next.Filter, next.second[i].key)
				}
			}
			*table = next
			return
		}
		seed += 2
		attempts++
		if attempts == 4 {
			capacity *= 2
			attempts = 0
		}
	}
}

func AddCuckoo(table *Cuckoo, key string) {
	if ContainsCuckoo(table, key) {
		return
	}
	if table.Size+1 > table.Capacity*4/5 {
		rebuild(table, table.Capacity*2, key)
		return
	}

	pending := key
	if placeKey(table, &pending) {
		table.Size++
		addToBloomFilter(&table.Filter, key)
	} else {
		rebuild(table, table.Capacity, pending)
	}
}

func RemoveCuckoo(table *Cuckoo, key string) bool {
	if !MightContainBloomFilter(&table.Filter, key) {
		return false
	}
	cell := &table.first[hashfunctions.CuckooIndex(key, table.Capacity, table.seed)]
	if !cell.occupied || cell.key != key {
		cell = &table.second[hashfunctions.CuckooIndex(key, table.Capacity, table.seed+1)]
		if !cell.occupied || cell.key != key {
			return false
		}
	}

	*cell = CuckooCell{}
	table.Size--
	return true
}

func ClearCuckoo(table *Cuckoo) {
	*table = Cuckoo{}
}
