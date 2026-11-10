package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func processAsteroid(survivors *Stack, size int, direction string) error {
	if direction == "L" {
		size = -size
	}

	for !isEmpty(survivors) && size < 0 {
		head, err := strconv.Atoi(survivors.head.value)
		if err != nil {
			return err
		}

		if head < 0 {
			break
		}

		size = head + size
		pop(survivors)
	}

	if size != 0 {
		push(survivors, strconv.Itoa(size))
	}

	return nil
}

func readAsteroids(survivors *Stack) error {
	input := bufio.NewReader(os.Stdin)

	var count int

	if _, err := fmt.Fscan(input, &count); err != nil {
		return err
	}

	if count < 0 {
		return os.ErrInvalid
	}

	for i := 0; i < count; i++ {
		var size int
		var direction string

		if _, err := fmt.Fscan(input, &size, &direction); err != nil {
			return err
		}

		if size <= 0 || (direction != "L" && direction != "R") {
			return os.ErrInvalid
		}

		if err := processAsteroid(survivors, size, direction); err != nil {
			return err
		}
	}

	return nil
}

func printAsteroids(survivors *Stack) error {
	var reversed Stack
	initStack(&reversed)
	defer clearStack(&reversed)

	for !isEmpty(survivors) {
		push(&reversed, survivors.head.value)
		pop(survivors)
	}

	fmt.Print("[")

	first := true

	for !isEmpty(&reversed) {
		value, err := strconv.Atoi(reversed.head.value)
		if err != nil {
			return err
		}

		size := value
		direction := "R"

		if value < 0 {
			size = -value
			direction = "L"
		}

		if !first {
			fmt.Print(", ")
		}

		fmt.Printf("(%d, %s)", size, direction)
		first = false

		pop(&reversed)
	}

	fmt.Println("]")
	return nil
}

func clearStack(stack *Stack) {
	for !isEmpty(stack) {
		pop(stack)
	}
}

func run() error {
	var survivors Stack
	initStack(&survivors)
	defer clearStack(&survivors)

	if err := readAsteroids(&survivors); err != nil {
		return err
	}

	return printAsteroids(&survivors)
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}