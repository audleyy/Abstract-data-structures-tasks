package hashtable

import "lab2/hash-functions/Golang"

type CuckooCell struct {
	key      string
	occupied bool
}

type Cuckoo struct {
	first    []CuckooCell
	second   []CuckooCell
	capacity int
	size     int
	seed     int
}

func initCuckooCapacity(table *Cuckoo, capacity int, seed int) {
	*table = Cuckoo{}
	table.first = make([]CuckooCell, capacity)
	table.second = make([]CuckooCell, capacity)
	table.capacity = capacity
	table.seed = seed
}

func InitCuckoo(table *Cuckoo) {
	initCuckooCapacity(table, 2, 1)
}

func ContainsCuckoo(table *Cuckoo, key string) bool {
	first := table.first[hashfunctions.CuckooIndex(key, table.capacity, table.seed)]
	if first.occupied && first.key == key {
		return true
	}

	second := table.second[hashfunctions.CuckooIndex(key, table.capacity, table.seed+1)]
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
		cell := &cells[hashfunctions.CuckooIndex(*key, table.capacity, seed)]
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
	for i := 0; i < source.capacity; i++ {
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
			next.size = table.size + 1
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
	if table.size+1 > table.capacity*4/5 {
		rebuild(table, table.capacity*2, key)
		return
	}

	pending := key
	if placeKey(table, &pending) {
		table.size++
	} else {
		rebuild(table, table.capacity, pending)
	}
}

func RemoveCuckoo(table *Cuckoo, key string) bool {
	cell := &table.first[hashfunctions.CuckooIndex(key, table.capacity, table.seed)]
	if !cell.occupied || cell.key != key {
		cell = &table.second[hashfunctions.CuckooIndex(key, table.capacity, table.seed+1)]
		if !cell.occupied || cell.key != key {
			return false
		}
	}

	*cell = CuckooCell{}
	table.size--
	return true
}

func ClearCuckoo(table *Cuckoo) {
	*table = Cuckoo{}
}
