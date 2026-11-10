package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func buildFrequencyList(text string, head **ListNode) {
	length := len(text)

	for i := 0; i < length; i++ {
		symbol := text[i : i+1]
		found := findNode(*head, symbol)

		if found != nil {
			found.value.frequency++
		} else {
			addToTail(head, createLeaf(symbol))
		}
	}
}

func extractMin(head **ListNode) *Node {
	if *head == nil {
		return nil
	}

	minimum := *head
	current := (*head).next

	for current != nil {
		if current.value.frequency < minimum.value.frequency {
			minimum = current
		}
		current = current.next
	}

	result := minimum.value
	deleteCurrent(head, minimum)
	return result
}

func buildTree(head **ListNode) *Node {
	for *head != nil && (*head).next != nil {
		left := extractMin(head)
		right := extractMin(head)
		parent := createParent(left, right)
		addToTail(head, parent)
	}

	return extractMin(head)
}

func findCode(root *Node, symbol string, path string, code *string) bool {
	if isLeaf(root) {
		if root.symbol != symbol {
			return false
		}
		if path == "" {
			*code = "0"
		} else {
			*code = path
		}
		return true
	}

	if findCode(root.left, symbol, path+"0", code) {
		return true
	}
	return findCode(root.right, symbol, path+"1", code)
}

func encode(root *Node, text string, encoded *string) {
	*encoded = ""
	length := len(text)

	for i := 0; i < length; i++ {
		var code string
		findCode(root, text[i:i+1], "", &code)
		if i != 0 {
			*encoded += " "
		}
		*encoded += code
	}
}

func decode(root *Node, encoded string, decoded *string) {
	*decoded = ""
	length := len(encoded)

	if isLeaf(root) {
		for i := 0; i < length; i++ {
			if encoded[i] != ' ' {
				*decoded += root.symbol
			}
		}
		return
	}

	current := root

	for i := 0; i < length; i++ {
		if encoded[i] == ' ' {
			continue
		}
		if encoded[i] == '0' {
			current = current.left
		} else {
			current = current.right
		}

		if isLeaf(current) {
			*decoded += current.symbol
			current = root
		}
	}
}

func main() {
	input := bufio.NewReader(os.Stdin)
	text, err := input.ReadString('\n')
	if err != nil && (err != io.EOF || len(text) == 0) {
		os.Exit(1)
	}
	if len(text) > 0 && text[len(text)-1] == '\n' {
		text = text[:len(text)-1]
	}
	if len(text) > 0 && text[len(text)-1] == '\r' {
		text = text[:len(text)-1]
	}

	var head *ListNode
	initForwardList(&head)
	buildFrequencyList(text, &head)
	root := buildTree(&head)

	var encoded string
	var decoded string
	encode(root, text, &encoded)
	decode(root, encoded, &decoded)

	length := len(text)
	for i := 0; i < length; i++ {
		var code string
		findCode(root, text[i:i+1], "", &code)
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%s:%s", text[i:i+1], code)
	}
	fmt.Println()
	fmt.Println(decoded)

	destroyTree(&root)
}
