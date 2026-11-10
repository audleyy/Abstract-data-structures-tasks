#include "array.h"

void init(Array& arr) {
	arr.data = nullptr;
	arr.size = 0;
}

void pushIndex(Array& arr, int index, const string& value) {
	if (index < 0 || index >= arr.size) return;

	string* newData = new string[arr.size + 1];
	
	for (int i = 0; i < index; ++i) {
		newData[i] = arr.data[i];
	}

	newData[index] = value;

	for (int j = index + 1; j < arr.size + 1; ++j) {
		newData[j] = arr.data[j - 1];
	}

	delete[] arr.data;
	arr.data = newData;
	arr.size++;
}

void pushBack(Array& arr, const string& value) {
	string* newData = new string[arr.size + 1];

	for (int i = 0; i < arr.size; ++i) {
		newData[i] = arr.data[i];
	}

	newData[arr.size] = value;

	delete[] arr.data;
	arr.data = newData;
	arr.size++;
}

string get(Array& arr, int index) {
	if (index < 0 || index >= arr.size) return string();
	return arr.data[index];
}

void remove(Array& arr, int index) {
	if (index < 0 || index >= arr.size) return;

	string* newData = new string[arr.size - 1];

	for (int i = 0, j = 0; i < arr.size; ++i) {
		if (i != index) {
			newData[j] = arr.data[i];
			j++;
		}
	}

	delete[] arr.data;
	arr.data = newData;
	arr.size--;
}

void replace(Array& arr, int index, const string& value) {
	if (index < 0 || index >= arr.size) return;

	arr.data[index] = value;
}

int length(Array& arr) {
	return arr.size;
}

void destroy(Array& arr) {
	delete[] arr.data;

	arr.data = nullptr;
	arr.size = 0;
}

void print(Array arr) {
	for (int i = 0; i < arr.size; ++i) {
		cout << arr.data[i] << ' ';
	}
	cout << endl;
}

