package hashtable

import "lab2/hash-functions/Golang"

type Node struct {
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

func Contains(set *Set, value string) bool {
	if !MightContainBloomFilter(&set.Filter, value) {
		return false
	}
	index := hashfunctions.HashFunction(value, set.Capacity)
	current := set.Buckets[index]

	for current != nil {
		if current.Value == value {
			return true
		}

		current = current.Next
	}

	return false
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

			newIndex := hashfunctions.HashFunction(current.Value, newCapacity)
			addToBloomFilter(&set.Filter, current.Value)

			current.Next = newBuckets[newIndex]
			newBuckets[newIndex] = current

			current = next
		}
	}

	set.Buckets = newBuckets
	set.Capacity = newCapacity
}

func Add(set *Set, value string) {
	if Contains(set, value) {
		return
	}

	if set.Size+1 > set.Capacity*3/4 {
		rehash(set)
	}

	index := hashfunctions.HashFunction(value, set.Capacity)

	newNode := &Node{
		Value: value,
		Next:  set.Buckets[index],
	}

	set.Buckets[index] = newNode
	set.Size++
	addToBloomFilter(&set.Filter, value)
}

func Remove(set *Set, value string) bool {
	if !MightContainBloomFilter(&set.Filter, value) {
		return false
	}
	index := hashfunctions.HashFunction(value, set.Capacity)
	current := set.Buckets[index]
	var previous *Node

	for current != nil {
		if current.Value == value {
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
