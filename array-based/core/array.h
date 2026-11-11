#ifndef ARRAY_H
#define ARRAY_H

#include <iostream>
#include <string>

using namespace std;

struct Array {
	string* data;
	int size;
};

void init(Array& arr);
void initFilledArray(Array& arr, int size, const string& value);
void pushIndex(Array& arr, int index, const string& value);
void pushBack(Array& arr, const string& value);
string get(const Array& arr, int index);
void remove(Array& arr, int index);
void replace(Array& arr, int index, const string& value);
int length(Array& arr);
void destroy(Array& arr);
void print(Array arr);


#endif
