package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readArray(arr *Array, target *int) bool {
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
		if err != nil {
			return false
		}
		pushBack(arr, strconv.Itoa(value))
	}

	if _, err := fmt.Fscan(input, &token); err != nil {
		return false
	}
	value, err := strconv.Atoi(token)
	if err != nil {
		return false
	}
	*target = value
	return true
}

func subarrayToString(arr *Array, start int, end int) string {
	result := "["

	for i := start; i <= end; i++ {
		if i != start {
			result += ", "
		}
		result += get(arr, i)
	}

	return result + "]"
}

func findSubarrays(arr *Array, target int, result *string) bool {
	*result = ""

	for start := 0; start < arr.size; start++ {
		sum := 0

		for end := start; end < arr.size; end++ {
			value, _ := strconv.Atoi(get(arr, end))
			sum += value

			if sum == target {
				if *result != "" {
					*result += "\n"
				}
				*result += subarrayToString(arr, start, end)
			}
		}
	}

	if *result == "" {
		*result = "[]"
	}
	return true
}

func main() {
	var arr Array
	var target int
	result := 0
	var output string

	initArray(&arr)

	if !readArray(&arr, &target) {
		result = 1
	} else if !findSubarrays(&arr, target, &output) {
		result = 1
	} else {
		fmt.Println(output)
	}

	destroy(&arr)
	os.Exit(result)
}
