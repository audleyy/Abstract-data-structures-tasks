package main

type Node struct {
	symbol    string
	frequency int
	left      *Node
	right     *Node
}

func initTree(root **Node) {
	*root = nil
}

func createLeaf(symbol string) *Node {
	return &Node{symbol, 1, nil, nil}
}

func createParent(left *Node, right *Node) *Node {
	return &Node{"", left.frequency + right.frequency, left, right}
}

func isLeaf(node *Node) bool {
	return node != nil && node.left == nil && node.right == nil
}

func destroyTree(root **Node) {
	if *root == nil {
		return
	}

	destroyTree(&(*root).left)
	destroyTree(&(*root).right)
	*root = nil
}
