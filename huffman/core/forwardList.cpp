#include "forwardList.h"

void init(ListNode*& head) {
	head = nullptr;
}

void addToTail(ListNode*& head, Node* value) {
	ListNode* newNode = new ListNode{value, nullptr};
	ListNode* current = head;

	if (head == nullptr) {
		head = newNode;
		return;
	}

	while (current->next != nullptr) {
		current = current->next;
	}

	current->next = newNode;
}

ListNode* findNode(ListNode* head, char symbol) {
	ListNode* current = head;

	while (current != nullptr) {
		if (current->value->symbol == symbol) return current;
		current = current->next;
	}

	return nullptr;
}

void deleteHead(ListNode*& head) {
	if (head == nullptr) return;

	ListNode* toDelete = head;
	head = toDelete->next;
	delete toDelete;
}

void deleteCurrent(ListNode*& head, ListNode* nodeFound) {
	if (nodeFound == head) {
		deleteHead(head);
		return;
	}

	ListNode* current = head;
	while (current->next != nodeFound) {
		current = current->next;
	}

	current->next = nodeFound->next;
	delete nodeFound;
}

void clear(ListNode*& head) {
	while (head != nullptr) {
		destroyTree(head->value);
		deleteHead(head);
	}
}
