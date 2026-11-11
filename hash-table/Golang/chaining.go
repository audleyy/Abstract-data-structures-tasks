package hashtable

import "lab2/hash-functions/Golang"

type Node struct {
	value string
	next  *Node
}

type Set struct {
	buckets  []*Node
	capacity int
	size     int
}

func InitSet(set *Set) {
	set.capacity = 4
	set.size = 0
	set.buckets = make([]*Node, set.capacity)
}

func Contains(set *Set, value string) bool {
	index := hashfunctions.HashFunction(value, set.capacity)
	current := set.buckets[index]

	for current != nil {
		if current.value == value {
			return true
		}

		current = current.next
	}

	return false
}

func rehash(set *Set) {
	newCapacity := set.capacity * 2
	newBuckets := make([]*Node, newCapacity)

	for i := 0; i < set.capacity; i++ {
		current := set.buckets[i]

		for current != nil {
			next := current.next

			newIndex := hashfunctions.HashFunction(current.value, newCapacity)

			current.next = newBuckets[newIndex]
			newBuckets[newIndex] = current

			current = next
		}
	}

	set.buckets = newBuckets
	set.capacity = newCapacity
}

func Add(set *Set, value string) {
	if Contains(set, value) {
		return
	}

	if set.size+1 > set.capacity*3/4 {
		rehash(set)
	}

	index := hashfunctions.HashFunction(value, set.capacity)

	newNode := &Node{
		value: value,
		next:  set.buckets[index],
	}

	set.buckets[index] = newNode
	set.size++
}

func Remove(set *Set, value string) bool {
	index := hashfunctions.HashFunction(value, set.capacity)
	current := set.buckets[index]
	var previous *Node

	for current != nil {
		if current.value == value {
			if previous == nil {
				set.buckets[index] = current.next
			} else {
				previous.next = current.next
			}

			set.size--
			return true
		}

		previous = current
		current = current.next
	}

	return false
}

func Clear(set *Set) {
	set.buckets = nil
	set.capacity = 0
	set.size = 0
}
