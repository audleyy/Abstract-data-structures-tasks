package main

import "fmt"

type Node struct {
  value string
  next  *Node
}

type Stack struct {
  head *Node
}

func initStack(stack *Stack) {
  stack.head = nil
}

func isEmpty(stack *Stack) bool {
  return stack.head == nil
}

func push(stack *Stack, value string) {
  newNode := &Node{value, nil}
  newNode.next = stack.head
  stack.head = newNode
}

func pop(stack *Stack) {
  if isEmpty(stack) {
    return
  }
  toDelete := stack.head
  stack.head = toDelete.next
}

func print(stack *Stack) {
  current := stack.head
  for current != nil {
    fmt.Printf("%s ", current.value)
    current = current.next
  }
  fmt.Println()
}
