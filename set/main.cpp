#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <filesystem>

#include "core/set.h"

using namespace std;

struct Query {
	string command;
	string name;
	string value;
};

bool parseArguments(int argc, char* argv[], string& filename, string& queryText) {
	if (argc != 5) return false;

	if (string(argv[1]) != "--file" ||
		string(argv[3]) != "--query") {
		return false;
	}

	filename = argv[2];
	queryText = argv[4];

	return !filename.empty() && !queryText.empty();
}

bool parseQuery(const string& text, Query& query) {
	istringstream input(text);
	string extra;

	if (!(input >> query.command >> query.name)) {
		return false;
	}

	if (query.command == "PRINT") {
		return !(input >> extra);
	}

	if (query.command != "SETADD" &&
		query.command != "SETDEL" &&
		query.command != "SET_AT") {
		return false;
	}

	if (!(input >> query.value)) {
		return false;
	}

	return !(input >> extra);
}

bool readFile(const string& filename, vector<string>& lines) {
	error_code error;
	bool exists = filesystem::exists(filename, error);

	if (error) return false;

	if (!exists) return true;

	ifstream input(filename);
	if (!input) return false;

	string line;

	while (getline(input, line)) {
		lines.push_back(line);
	}

	return !input.bad();
}

bool findSet(const vector<string>& lines, const string& name, size_t& position) {
	position = lines.size();

	for (size_t i = 0; i < lines.size(); ++i) {
		istringstream row(lines[i]);
		string savedType;
		string savedName;

		if (!(row >> savedType >> savedName)) continue;
		if (savedName != name) continue;

		if (position != lines.size() || savedType != "E") {
			return false;
		}

		position = i;
	}

	return true;
}

void loadSet(const string& line, Set* set) {
	istringstream row(line);

	string type;
	string name;
	string value;

	row >> type >> name;

	while (row >> value) {
		add(set, value);
	}
}

string setToString(const Set* set, const string& name) {
	ostringstream output;

	output << "E " << name;

	for (size_t i = 0; i < set->table.capacity; ++i) {
		Node* current = set->table.buckets[i];

		while (current != nullptr) {
			output << ' ' << current->value;
			current = current->next;
		}
	}

	return output.str();
}

void printSet(const Set* set) {
	bool first = true;

	for (size_t i = 0; i < set->table.capacity; ++i) {
		Node* current = set->table.buckets[i];

		while (current != nullptr) {
			if (!first) cout << ' ';

			cout << current->value;
			first = false;
			current = current->next;
		}
	}

	cout << '\n';
}

bool writeFile(const string& filename, const vector<string>& lines) {
	ofstream output(filename);
	if (!output) return false;

	for (const string& line : lines) {
		output << line << '\n';
	}

	output.close();
	return !output.fail();
}

int main(int argc, char* argv[]) {
	string filename;
	string queryText;
	Query query;

	if (!parseArguments(argc, argv, filename, queryText)) {
		return 1;
	}

	if (!parseQuery(queryText, query)) {
		return 1;
	}

	vector<string> lines;

	if (!readFile(filename, lines)) {
		return 1;
	}

	size_t position;

	if (!findSet(lines, query.name, position)) {
		return 1;
	}

	bool exists = position != lines.size();

	if (!exists && query.command != "SETADD") {
		return 1;
	}

	Set currentSet;
	init(&currentSet);

	if (exists) {
		loadSet(lines[position], &currentSet);
	}

	bool changed = false;

	if (query.command == "SETADD") {
		int oldSize = currentSet.table.size;
		add(&currentSet, query.value);
		changed = currentSet.table.size != oldSize;
	} else if (query.command == "SETDEL") {
		changed = remove(&currentSet, query.value);
	} else if (query.command == "SET_AT") {
		cout << (contains(&currentSet, query.value) ? "TRUE\n" : "FALSE\n");
	} else if (query.command == "PRINT") {
		printSet(&currentSet);
	}

	if (changed) {
		string updated = setToString(&currentSet, query.name);

		if (exists) {
			lines[position] = updated;
		} else {
			lines.push_back(updated);
		}

		if (!writeFile(filename, lines)) {
			clear(&currentSet);
			return 1;
		}
	}

	clear(&currentSet);

	if (query.command == "SETADD" || query.command == "SETDEL") {
		cout << (changed ? "OK\n" : "Ничего не изменилось\n");
	}

	return 0;
}