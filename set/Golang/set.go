package main

type Node struct {
	value string
	next  *Node
}

type Set struct {
	buckets  []*Node
	capacity int
	size     int
}

func initSet(set *Set) {
	set.capacity = 4
	set.size = 0
	set.buckets = make([]*Node, set.capacity)
}

func hashFunction(value string, capacity int) int {
	var hash uint

	for i := 0; i < len(value); i++ {
		hash = hash*31 + uint(value[i])
	}

	return int(hash % uint(capacity))
}

func contains(set *Set, value string) bool {
	index := hashFunction(value, set.capacity)
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

			newIndex := hashFunction(current.value, newCapacity)

			current.next = newBuckets[newIndex]
			newBuckets[newIndex] = current

			current = next
		}
	}

	set.buckets = newBuckets
	set.capacity = newCapacity
}

func add(set *Set, value string) {
	if contains(set, value) {
		return
	}

	if set.size+1 > set.capacity*3/4 {
		rehash(set)
	}

	index := hashFunction(value, set.capacity)

	newNode := &Node{
		value: value,
		next:  set.buckets[index],
	}

	set.buckets[index] = newNode
	set.size++
}

func remove(set *Set, value string) bool {
	index := hashFunction(value, set.capacity)
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

func clear(set *Set) {
	set.buckets = nil
	set.capacity = 0
	set.size = 0
}