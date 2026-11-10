#include "stack.h"

void init(Stack* stack) {
	stack->head = nullptr;
}

bool isEmpty(Stack* stack) {
	return stack->head == nullptr;
}

void push(Stack* stack, const string& value) {
	Node* newNode = new Node{value, nullptr};

	newNode->next = stack->head;
	stack->head = newNode;
}

void pop(Stack* stack) {
	if (isEmpty(stack)) return;

	Node* toDelete = stack->head;

	stack->head = toDelete->next;
	delete toDelete;
}

void print(Stack* stack) {
	Node* current = stack->head;

	while (current != nullptr) {
		cout << current->value << " ";
		current = current->next;
	}

	cout << endl;
}
