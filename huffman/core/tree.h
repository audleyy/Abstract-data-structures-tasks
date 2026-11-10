#ifndef TREE_H
#define TREE_H

struct Node {
	char symbol;
	int frequency;
	Node* left;
	Node* right;
};

void init(Node*& root);
Node* createLeaf(char symbol);
Node* createParent(Node* left, Node* right);
bool isLeaf(const Node* node);
void destroyTree(Node*& root);

#endif
