package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"lab2/hash-table/Golang"
)

func copySet(source *hashtable.Set, target *hashtable.Set) {
	hashtable.Clear(target)
	hashtable.InitSet(target)

	capacity := source.Capacity
	for i := 0; i < capacity; i++ {
		current := source.Buckets[i]

		for current != nil {
			hashtable.Add(target, current.Key)
			current = current.Next
		}
	}
}

func findPartition(remaining *hashtable.Set, current *hashtable.Set, best *hashtable.Set, sum int32, half int32, bestSum *int32) {
	if sum > *bestSum {
		copySet(current, best)
		*bestSum = sum
	}

	if *bestSum == half || remaining.Size == 0 {
		return
	}

	var value string
	capacity := remaining.Capacity
	for i := 0; i < capacity; i++ {
		if remaining.Buckets[i] != nil {
			value = remaining.Buckets[i].Key
			break
		}
	}

	number, _ := strconv.Atoi(value)
	hashtable.Remove(remaining, value)

	if int32(number) <= half-sum {
		hashtable.Add(current, value)
		findPartition(remaining, current, best, sum+int32(number), half, bestSum)
		hashtable.Remove(current, value)
	}

	if *bestSum != half {
		findPartition(remaining, current, best, sum, half, bestSum)
	}

	hashtable.Add(remaining, value)
}

func partitionSet(source *hashtable.Set, first *hashtable.Set, second *hashtable.Set) int32 {
	var total int32
	capacity := source.Capacity

	for i := 0; i < capacity; i++ {
		node := source.Buckets[i]

		for node != nil {
			value, err := strconv.Atoi(node.Key)
			if err != nil || value <= 0 {
				return -1
			}

			total += int32(value)
			node = node.Next
		}
	}

	var remaining hashtable.Set
	var current hashtable.Set
	var best hashtable.Set
	var bestSum int32

	hashtable.InitSet(&remaining)
	hashtable.InitSet(&current)
	hashtable.InitSet(&best)
	copySet(source, &remaining)
	findPartition(&remaining, &current, &best, 0, total/2, &bestSum)
	copySet(&best, first)

	for i := 0; i < capacity; i++ {
		node := source.Buckets[i]

		for node != nil {
			if !hashtable.Contains(first, node.Key) {
				hashtable.Add(second, node.Key)
			}
			node = node.Next
		}
	}

	hashtable.Clear(&remaining)
	hashtable.Clear(&current)
	hashtable.Clear(&best)
	return total - 2*bestSum
}

func readSet(set *hashtable.Set) bool {
	input := bufio.NewReader(os.Stdin)
	var token string
	if _, err := fmt.Fscan(input, &token); err != nil {
		return false
	}
	count, err := strconv.ParseInt(token, 10, 32)
	if err != nil || count < 0 {
		return false
	}

	for i := int64(0); i < count; i++ {
		if _, err := fmt.Fscan(input, &token); err != nil {
			return false
		}
		value, err := strconv.ParseInt(token, 10, 32)
		if err != nil || value <= 0 {
			return false
		}
		hashtable.Add(set, strconv.Itoa(int(value)))
	}

	return true
}

func setToString(set *hashtable.Set) string {
	result := "{"
	first := true

	for i := 0; i < set.Capacity; i++ {
		current := set.Buckets[i]

		for current != nil {
			if !first {
				result += ", "
			}
			result += current.Key
			first = false
			current = current.Next
		}
	}

	return result + "}"
}

func main() {
	var source hashtable.Set
	var first hashtable.Set
	var second hashtable.Set
	result := 0

	hashtable.InitSet(&source)
	hashtable.InitSet(&first)
	hashtable.InitSet(&second)

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

	hashtable.Clear(&source)
	hashtable.Clear(&first)
	hashtable.Clear(&second)

	os.Exit(result)
}
