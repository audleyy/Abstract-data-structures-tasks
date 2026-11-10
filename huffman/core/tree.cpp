#include "tree.h"

void init(Node*& root) {
	root = nullptr;
}

Node* createLeaf(char symbol) {
	return new Node{symbol, 1, nullptr, nullptr};
}

Node* createParent(Node* left, Node* right) {
	return new Node{'\0', left->frequency + right->frequency, left, right};
}

bool isLeaf(const Node* node) {
	return node != nullptr && node->left == nullptr && node->right == nullptr;
}

void destroyTree(Node*& root) {
	if (root == nullptr) return;

	destroyTree(root->left);
	destroyTree(root->right);
	delete root;
	root = nullptr;
}
