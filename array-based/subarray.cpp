#include "core/array.h"

bool readArray(Array& arr, int& target) {
	int count;
	if (!(cin >> count) || count < 0) return false;

	for (int i = 0; i < count; ++i) {
		int value;
		if (!(cin >> value)) return false;
		pushBack(arr, to_string(value));
	}

	if (!(cin >> target)) return false;
	return true;
}

string subarrayToString(Array& arr, int start, int end) {
	string result = "[";

	for (int i = start; i <= end; ++i) {
		if (i != start) result += ", ";
		result += get(arr, i);
	}

	return result + "]";
}

bool findSubarrays(Array& arr, int target, string& result) {
	result = "";

	for (int start = 0; start < arr.size; ++start) {
		int sum = 0;

		for (int end = start; end < arr.size; ++end) {
			int value = stoi(get(arr, end));
			sum += value;

			if (sum == target) {
				if (!result.empty()) result += '\n';
				result += subarrayToString(arr, start, end);
			}
		}
	}

	if (result.empty()) result = "[]";
	return true;
}

int main() {
	Array arr{nullptr, 0};
	int target;
	int result = 0;
	string output;

	init(arr);

	if (!readArray(arr, target)) {
		result = 1;
	} else if (!findSubarrays(arr, target, output)) {
		result = 1;
	} else {
		cout << output << '\n';
	}

	destroy(arr);
	return result;
}
