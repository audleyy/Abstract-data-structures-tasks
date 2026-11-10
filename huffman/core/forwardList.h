#ifndef FORWARD_LIST_H
#define FORWARD_LIST_H

#include "tree.h"

struct ListNode {
	Node* value;
	ListNode* next;
};

void init(ListNode*& head);
void addToTail(ListNode*& head, Node* value);
ListNode* findNode(ListNode* head, char symbol);
void deleteHead(ListNode*& head);
void deleteCurrent(ListNode*& head, ListNode* nodeFound);
void clear(ListNode*& head);

#endif
