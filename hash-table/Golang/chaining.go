package hashtable

import "lab2/hash-functions/Golang"

type Node struct {
	Key   string
	Value string
	Next  *Node
}

type Set struct {
	Filter   BloomFilter
	Buckets  []*Node
	Capacity int
	Size     int
}

func InitSet(set *Set) {
	set.Capacity = 4
	set.Size = 0
	set.Buckets = make([]*Node, set.Capacity)
	initBloomFilter(&set.Filter, set.Capacity*16)
}

func findNode(set *Set, key string) *Node {
	if !MightContainBloomFilter(&set.Filter, key) {
		return nil
	}
	index := hashfunctions.HashFunction(key, set.Capacity)
	current := set.Buckets[index]

	for current != nil {
		if current.Key == key {
			return current
		}

		current = current.Next
	}

	return nil
}

func Contains(set *Set, key string) bool {
	return findNode(set, key) != nil
}

func Get(set *Set, key string) (string, bool) {
	node := findNode(set, key)
	if node == nil {
		return "", false
	}
	return node.Value, true
}

func rehash(set *Set) {
	newCapacity := set.Capacity * 2
	newBuckets := make([]*Node, newCapacity)
	clearBloomFilter(&set.Filter)
	initBloomFilter(&set.Filter, newCapacity*16)

	for i := 0; i < set.Capacity; i++ {
		current := set.Buckets[i]

		for current != nil {
			next := current.Next

			newIndex := hashfunctions.HashFunction(current.Key, newCapacity)
			addToBloomFilter(&set.Filter, current.Key)

			current.Next = newBuckets[newIndex]
			newBuckets[newIndex] = current

			current = next
		}
	}

	set.Buckets = newBuckets
	set.Capacity = newCapacity
}

func insert(set *Set, key string, value string) {
	if set.Size+1 > set.Capacity*3/4 {
		rehash(set)
	}

	index := hashfunctions.HashFunction(key, set.Capacity)

	newNode := &Node{
		Key:   key,
		Value: value,
		Next:  set.Buckets[index],
	}

	set.Buckets[index] = newNode
	set.Size++
	addToBloomFilter(&set.Filter, key)
}

func Add(set *Set, key string) {
	if !Contains(set, key) {
		insert(set, key, "")
	}
}

func Put(set *Set, key string, value string) {
	node := findNode(set, key)
	if node != nil {
		node.Value = value
		return
	}
	insert(set, key, value)
}

func Remove(set *Set, key string) bool {
	if !MightContainBloomFilter(&set.Filter, key) {
		return false
	}
	index := hashfunctions.HashFunction(key, set.Capacity)
	current := set.Buckets[index]
	var previous *Node

	for current != nil {
		if current.Key == key {
			if previous == nil {
				set.Buckets[index] = current.Next
			} else {
				previous.Next = current.Next
			}

			set.Size--
			return true
		}

		previous = current
		current = current.Next
	}

	return false
}

func Clear(set *Set) {
	clearBloomFilter(&set.Filter)
	set.Buckets = nil
	set.Capacity = 0
	set.Size = 0
}
