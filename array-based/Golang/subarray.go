package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readArray(arr *Array, target *int32) bool {
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
		if err != nil {
			return false
		}
		pushBack(arr, strconv.Itoa(int(value)))
	}

	if _, err := fmt.Fscan(input, &token); err != nil {
		return false
	}
	value, err := strconv.ParseInt(token, 10, 32)
	if err != nil {
		return false
	}
	*target = int32(value)
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

func findSubarrays(arr *Array, target int32, result *string) bool {
	*result = ""

	for start := 0; start < arr.size; start++ {
		var sum int32

		for end := start; end < arr.size; end++ {
			value, _ := strconv.Atoi(get(arr, end))
			sum += int32(value)

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
	var target int32
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
