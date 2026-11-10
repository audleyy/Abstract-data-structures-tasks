#ifndef STACK_H
#define STACK_H

#include <iostream>
#include <string>

using namespace std;

struct Node {
	string value;
	Node* next;
};

struct Stack {
	Node* head;
};

void init(Stack* stack);
bool isEmpty(Stack* stack);
void push(Stack* stack, const string& value);
void pop(Stack* stack);
void print(Stack* stack);


#endif
