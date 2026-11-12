package hashtable

import "lab2/hash-functions/Golang"

type CuckooCell struct {
	key      string
	value    string
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

func findCell(table *Cuckoo, key string) *CuckooCell {
	if !MightContainBloomFilter(&table.Filter, key) {
		return nil
	}
	first := &table.first[hashfunctions.CuckooIndex(key, table.Capacity, table.seed)]
	if first.occupied && first.key == key {
		return first
	}

	second := &table.second[hashfunctions.CuckooIndex(key, table.Capacity, table.seed+1)]
	if second.occupied && second.key == key {
		return second
	}
	return nil
}

func ContainsCuckoo(table *Cuckoo, key string) bool {
	return findCell(table, key) != nil
}

func GetCuckoo(table *Cuckoo, key string) (string, bool) {
	cell := findCell(table, key)
	if cell == nil {
		return "", false
	}
	return cell.value, true
}

func placeEntry(table *Cuckoo, entry *CuckooCell) bool {
	first := true
	for step := 0; step < 32; step++ {
		cells := table.first
		seed := table.seed
		if !first {
			cells = table.second
			seed++
		}
		cell := &cells[hashfunctions.CuckooIndex(entry.key, table.Capacity, seed)]
		*cell, *entry = *entry, *cell
		if !entry.occupied {
			return true
		}
		first = !first
	}

	return false
}

func copyEntries(source *Cuckoo, target *Cuckoo) bool {
	for i := 0; i < source.Capacity; i++ {
		if source.first[i].occupied {
			entry := source.first[i]
			if !placeEntry(target, &entry) {
				return false
			}
		}
		if source.second[i].occupied {
			entry := source.second[i]
			if !placeEntry(target, &entry) {
				return false
			}
		}
	}
	return true
}

func rebuild(table *Cuckoo, capacity int, extra CuckooCell) {
	seed := table.seed + 2
	attempts := 0

	for {
		var next Cuckoo
		initCuckooCapacity(&next, capacity, seed)
		entry := extra
		placed := copyEntries(table, &next) && placeEntry(&next, &entry)
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

func insertCuckoo(table *Cuckoo, key string, value string) {
	pending := CuckooCell{key: key, value: value, occupied: true}
	if table.Size+1 > table.Capacity*4/5 {
		rebuild(table, table.Capacity*2, pending)
		return
	}

	if placeEntry(table, &pending) {
		table.Size++
		addToBloomFilter(&table.Filter, key)
	} else {
		rebuild(table, table.Capacity, pending)
	}
}

func AddCuckoo(table *Cuckoo, key string) {
	if !ContainsCuckoo(table, key) {
		insertCuckoo(table, key, "")
	}
}

func PutCuckoo(table *Cuckoo, key string, value string) {
	cell := findCell(table, key)
	if cell != nil {
		cell.value = value
		return
	}
	insertCuckoo(table, key, value)
}

func RemoveCuckoo(table *Cuckoo, key string) bool {
	cell := findCell(table, key)
	if cell == nil {
		return false
	}

	*cell = CuckooCell{}
	table.Size--
	return true
}

func ClearCuckoo(table *Cuckoo) {
	*table = Cuckoo{}
}
