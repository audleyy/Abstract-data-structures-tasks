#include "../set/core/set.h"

#include <iostream>

#include <climits>

void copySet(const Set* source, Set* target) {
	clear(target);
	init(target);

	int capacity = source->table.capacity;
	for (int i = 0; i < capacity; ++i) {
		Node* current = source->table.buckets[i];

		while (current != nullptr) {
			add(target, current->data.first);
			current = current->next;
		}
	}
}

void findPartition(Set* remaining, Set* current, Set* best, int sum, int half, int& bestSum) {
	if (sum > bestSum) {
		copySet(current, best);
		bestSum = sum;
	}

	if (bestSum == half || remaining->table.size == 0) return;

	string value;
	int capacity = remaining->table.capacity;
	for (int i = 0; i < capacity; ++i) {
		if (remaining->table.buckets[i] != nullptr) {
			value = remaining->table.buckets[i]->data.first;
			break;
		}
	}

	int number = stoi(value);
	remove(remaining, value);

	if (number <= half - sum) {
		add(current, value);
		findPartition(remaining, current, best, sum + number, half, bestSum);
		remove(current, value);
	}

	if (bestSum != half) {
		findPartition(remaining, current, best, sum, half, bestSum);
	}

	add(remaining, value);
}

int partitionSet(const Set* source, Set* first, Set* second) {
	int total = 0;
	int capacity = source->table.capacity;

	for (int i = 0; i < capacity; ++i) {
		Node* node = source->table.buckets[i];

		while (node != nullptr) {
			int value = stoi(node->data.first);
			if (value <= 0) return -1;

			total += value;
			node = node->next;
		}
	}

	Set remaining{};
	Set current{};
	Set best{};
	int bestSum = 0;

	init(&remaining);
	init(&current);
	init(&best);
	copySet(source, &remaining);
	findPartition(&remaining, &current, &best, 0, total / 2, bestSum);
	copySet(&best, first);

	for (int i = 0; i < capacity; ++i) {
		Node* node = source->table.buckets[i];

		while (node != nullptr) {
			if (!contains(first, node->data.first)) add(second, node->data.first);
			node = node->next;
		}
	}

	clear(&remaining);
	clear(&current);
	clear(&best);
	return total - 2 * bestSum;
}

bool readSet(Set* set) {
	int count;
	if (!(cin >> count) || count < 0) return false;

	for (int i = 0; i < count; ++i) {
		int value;
		if (!(cin >> value) || value <= 0) return false;
		add(set, to_string(value));
	}

	return true;
}

string setToString(const Set* set) {
	string result = "{";
	bool first = true;

	for (size_t i = 0; i < set->table.capacity; ++i) {
		Node* current = set->table.buckets[i];

		while (current != nullptr) {
			if (!first) result += ", ";
			result += current->data.first;
			first = false;
			current = current->next;
		}
	}

	return result + "}";
}

int main() {
	Set source{};
	Set first{};
	Set second{};
	int result = 0;

	init(&source);
	init(&first);
	init(&second);

	if (!readSet(&source)) {
		result = 1;
	} else {
		int difference = partitionSet(&source, &first, &second);

		if (difference == -1) {
			result = 1;
		} else {
			string firstText = setToString(&first);
			string secondText = setToString(&second);

			cout << firstText << '\n';
			cout << secondText << '\n';
			cout << difference << '\n';
		}
	}

	clear(&source);
	clear(&first);
	clear(&second);
	return result;
}
