package main

import (
	"bufio"
	"os"
	"fmt"
	"math"
	"strconv"
)

func copySet(source *Set, target *Set) {
	clear(target)
	initSet(target)

	capacity := source.capacity
	for i := 0; i < capacity; i++ {
		current := source.buckets[i]

		for current != nil {
			add(target, current.value)
			current = current.next
		}
	}
}

func findPartition(remaining *Set, current *Set, best *Set, sum int, half int, bestSum *int) {
	if sum > *bestSum {
		copySet(current, best)
		*bestSum = sum
	}

	if *bestSum == half || remaining.size == 0 {
		return
	}

	var value string
	capacity := remaining.capacity
	for i := 0; i < capacity; i++ {
		if remaining.buckets[i] != nil {
			value = remaining.buckets[i].value
			break
		}
	}

	number, _ := strconv.Atoi(value)
	remove(remaining, value)

	if number <= half-sum {
		add(current, value)
		findPartition(remaining, current, best, sum+number, half, bestSum)
		remove(current, value)
	}

	if *bestSum != half {
		findPartition(remaining, current, best, sum, half, bestSum)
	}

	add(remaining, value)
}

func partitionSet(source *Set, first *Set, second *Set) int {
	total := 0
	capacity := source.capacity

	for i := 0; i < capacity; i++ {
		node := source.buckets[i]

		for node != nil {
			value, err := strconv.Atoi(node.value)
			if err != nil || value <= 0 || value > math.MaxInt-total {
				return -1
			}

			total += value
			node = node.next
		}
	}

	var remaining Set
	var current Set
	var best Set
	bestSum := 0

	initSet(&remaining)
	initSet(&current)
	initSet(&best)
	copySet(source, &remaining)
	findPartition(&remaining, &current, &best, 0, total/2, &bestSum)
	copySet(&best, first)

	for i := 0; i < capacity; i++ {
		node := source.buckets[i]

		for node != nil {
			if !contains(first, node.value) {
				add(second, node.value)
			}
			node = node.next
		}
	}

	clear(&remaining)
	clear(&current)
	clear(&best)
	return total - 2*bestSum
}

func readSet(set *Set) bool {
	input := bufio.NewReader(os.Stdin)
	var token string
	if _, err := fmt.Fscan(input, &token); err != nil {
		return false
	}
	count, err := strconv.Atoi(token)
	if err != nil || count < 0 {
		return false
	}

	for i := 0; i < count; i++ {
		if _, err := fmt.Fscan(input, &token); err != nil {
			return false
		}
		value, err := strconv.Atoi(token)
		if err != nil || value <= 0 {
			return false
		}
		add(set, strconv.Itoa(value))
	}

	return true
}

func setToString(set *Set) string {
	result := "{"
	first := true

	for i := 0; i < set.capacity; i++ {
		current := set.buckets[i]

		for current != nil {
			if !first {
				result += ", "
			}
			result += current.value
			first = false
			current = current.next
		}
	}

	return result + "}"
}

func main() {
	var source Set
	var first Set
	var second Set
	result := 0

	initSet(&source)
	initSet(&first)
	initSet(&second)

	if !readSet(&source) {
		result = 1
	} else {
		difference := partitionSet(&source, &first, &second)

		if difference == -1 {
			result = 1
		} else {
			firstText := setToString(&first)
			secondText := setToString(&second)

			fmt.Println(firstText)
			fmt.Println(secondText)
			fmt.Println(difference)
		}
	}

	clear(&source)
	clear(&first)
	clear(&second)

	os.Exit(result)
}
