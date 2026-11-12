package main

import (
	"fmt"
	"os"
	"strings"

	"lab2/hash-table/Golang"
)

type Query struct {
	command string
	name    string
	value   string
}

func parseArguments(args []string) (string, string, error) {
	if len(args) != 5 ||
		args[1] != "--file" ||
		args[3] != "--query" {
		return "", "", os.ErrInvalid
	}

	filename := args[2]
	queryText := args[4]

	if filename == "" || queryText == "" {
		return "", "", os.ErrInvalid
	}

	return filename, queryText, nil
}

func splitFields(text string) []string {
	return strings.FieldsFunc(text, func(symbol rune) bool {
		return symbol == ' ' || (symbol >= '\t' && symbol <= '\r')
	})
}

func parseQuery(text string) (Query, error) {
	var query Query
	fields := splitFields(text)

	if len(fields) < 2 {
		return query, os.ErrInvalid
	}

	query.command = fields[0]
	query.name = fields[1]

	switch query.command {
	case "SETADD", "SETDEL", "SET_AT":
		if len(fields) != 3 {
			return query, os.ErrInvalid
		}

		query.value = fields[2]

	case "PRINT":
		if len(fields) != 2 {
			return query, os.ErrInvalid
		}

	default:
		return query, os.ErrInvalid
	}

	return query, nil
}

func readFile(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)

	if os.IsNotExist(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	text := strings.TrimSuffix(string(data), "\n")
	return strings.Split(text, "\n"), nil
}

func findSet(lines []string, name string) (int, error) {
	position := -1

	for i, line := range lines {
		fields := splitFields(line)

		if len(fields) < 2 || fields[1] != name {
			continue
		}

		if fields[0] != "E" || position != -1 {
			return -1, os.ErrInvalid
		}

		position = i
	}

	return position, nil
}

func loadSet(line string, set *hashtable.Set) {
	fields := splitFields(line)

	for _, value := range fields[2:] {
		hashtable.Add(set, value)
	}
}

func setValues(set *hashtable.Set) []string {
	var values []string

	for i := 0; i < set.Capacity; i++ {
		current := set.Buckets[i]

		for current != nil {
			values = append(values, current.Value)
			current = current.Next
		}
	}

	return values
}

func setToString(set *hashtable.Set, name string) string {
	result := "E " + name
	values := setValues(set)

	if len(values) > 0 {
		result += " " + strings.Join(values, " ")
	}

	return result
}

func printSet(set *hashtable.Set) {
	fmt.Println(strings.Join(setValues(set), " "))
}

func executeSet(query Query, set *hashtable.Set) (bool, error) {
	switch query.command {
	case "SETADD":
		oldSize := set.Size
		hashtable.Add(set, query.value)
		return set.Size != oldSize, nil

	case "SETDEL":
		return hashtable.Remove(set, query.value), nil

	case "SET_AT":
		if hashtable.Contains(set, query.value) {
			fmt.Println("TRUE")
		} else {
			fmt.Println("FALSE")
		}
		return false, nil

	case "PRINT":
		printSet(set)
		return false, nil
	}

	return false, os.ErrInvalid
}

func writeFile(filename string, lines []string) error {
	content := strings.Join(lines, "\n") + "\n"
	return os.WriteFile(filename, []byte(content), 0666)
}

func run() error {
	filename, queryText, err := parseArguments(os.Args)
	if err != nil {
		return err
	}

	query, err := parseQuery(queryText)
	if err != nil {
		return err
	}

	lines, err := readFile(filename)
	if err != nil {
		return err
	}

	position, err := findSet(lines, query.name)
	if err != nil {
		return err
	}

	if position == -1 && query.command != "SETADD" {
		return os.ErrNotExist
	}

	var set hashtable.Set
	hashtable.InitSet(&set)
	defer hashtable.Clear(&set)

	if position != -1 {
		loadSet(lines[position], &set)
	}

	changed, err := executeSet(query, &set)
	if err != nil {
		return err
	}

	if changed {
		updated := setToString(&set, query.name)

		if position == -1 {
			lines = append(lines, updated)
		} else {
			lines[position] = updated
		}

		if err := writeFile(filename, lines); err != nil {
			return err
		}
	}

	if query.command == "SETADD" || query.command == "SETDEL" {
		if changed {
			fmt.Println("OK")
		} else {
			fmt.Println("Ничего не изменилось")
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}
