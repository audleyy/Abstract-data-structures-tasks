package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func run() int {
	input := bufio.NewReader(os.Stdin)
	var capacityText, countText string
	if _, err := fmt.Fscan(input, &capacityText, &countText); err != nil {
		return 1
	}
	capacity, capacityErr := strconv.ParseInt(capacityText, 10, 32)
	queryCount, countErr := strconv.ParseInt(countText, 10, 32)
	if capacityErr != nil || countErr != nil || capacity < 0 || queryCount < 0 {
		return 1
	}

	cache := newLFUCache(int(capacity))
	firstResult := true

	for i := int64(0); i < queryCount; i++ {
		var command, keyText string
		if _, err := fmt.Fscan(input, &command, &keyText); err != nil {
			return 1
		}
		key, err := strconv.ParseInt(keyText, 10, 32)
		if err != nil {
			return 1
		}

		switch command {
		case "SET":
			var valueText string
			if _, err := fmt.Fscan(input, &valueText); err != nil {
				return 1
			}
			value, err := strconv.ParseInt(valueText, 10, 32)
			if err != nil {
				return 1
			}
			cache.set(int32(key), int32(value))
		case "GET":
			if !firstResult {
				fmt.Print(" ")
			}
			fmt.Print(cache.get(int32(key)))
			firstResult = false
		default:
			return 1
		}
	}

	fmt.Println()
	return 0
}

func main() {
	os.Exit(run())
}
