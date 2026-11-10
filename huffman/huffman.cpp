#include <iostream>
#include <string>

#include "core/tree.h"
#include "core/forwardList.h"

using namespace std;

void buildFrequencyList(const string& text, ListNode*& head) {
	int length = text.length();

	for (int i = 0; i < length; ++i) {
		ListNode* found = findNode(head, text[i]);

		if (found != nullptr) {
			found->value->frequency++;
		} else {
			addToTail(head, createLeaf(text[i]));
		}
	}
}

Node* extractMin(ListNode*& head) {
	if (head == nullptr) return nullptr;

	ListNode* minimum = head;
	ListNode* current = head->next;

	while (current != nullptr) {
		if (current->value->frequency < minimum->value->frequency) {
			minimum = current;
		}
		current = current->next;
	}

	Node* result = minimum->value;
	deleteCurrent(head, minimum);
	return result;
}

Node* buildTree(ListNode*& head) {
	while (head != nullptr && head->next != nullptr) {
		Node* left = extractMin(head);
		Node* right = extractMin(head);
		Node* parent = createParent(left, right);
		addToTail(head, parent);
	}

	return extractMin(head);
}

bool findCode(const Node* root, char symbol, string path, string& code) {
	if (isLeaf(root)) {
		if (root->symbol != symbol) return false;
		code = path.empty() ? "0" : path;
		return true;
	}

	if (findCode(root->left, symbol, path + "0", code)) return true;
	return findCode(root->right, symbol, path + "1", code);
}

void encode(const Node* root, const string& text, string& encoded) {
	encoded = "";
	int length = text.length();

	for (int i = 0; i < length; ++i) {
		string code;
		findCode(root, text[i], "", code);
		if (i != 0) encoded += ' ';
		encoded += code;
	}
}

void decode(const Node* root, const string& encoded, string& decoded) {
	decoded = "";
	int length = encoded.length();

	if (isLeaf(root)) {
		for (int i = 0; i < length; ++i) {
			if (encoded[i] != ' ') decoded += root->symbol;
		}
		return;
	}

	const Node* current = root;

	for (int i = 0; i < length; ++i) {
		if (encoded[i] == ' ') continue;
		if (encoded[i] == '0') {
			current = current->left;
		} else {
			current = current->right;
		}

		if (isLeaf(current)) {
			decoded += current->symbol;
			current = root;
		}
	}
}

int main() {
	string text;
	if (!getline(cin, text)) return 1;
	if (!text.empty() && text.back() == '\r') text.pop_back();

	ListNode* head;
	init(head);
	buildFrequencyList(text, head);
	Node* root = buildTree(head);

	string encoded;
	string decoded;
	encode(root, text, encoded);
	decode(root, encoded, decoded);

	int length = text.length();
	for (int i = 0; i < length; ++i) {
		string code;
		findCode(root, text[i], "", code);
		if (i != 0) cout << ' ';
		cout << text[i] << ':' << code;
	}
	cout << '\n';
	cout << decoded << '\n';

	destroyTree(root);
	return 0;
}
