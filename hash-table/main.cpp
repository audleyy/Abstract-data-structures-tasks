#include <iostream>
#include <vector>
#include <chrono>
#include <iomanip>
#include <utility>

#include "core/chaining.h"
#include "core/cuckoo.h"

using namespace std;
using namespace chrono;


void printBloomResults(const Chaining* chaining, const Cuckoo* cuckoo, const string& key) {
	cout << "Фильтр цепочек: " << (mightContainBloomFilter(&chaining->filter, key) ? "возможно присутствует" : "точно отсутствует") << '\n';

	cout << "Фильтр кукушки: " << (mightContainBloomFilter(&cuckoo->filter, key) ? "возможно присутствует" : "точно отсутствует") << '\n';
}

int main() {
	int count;
	if (!(cin >> count) || count < 0) {
		return 1;
	}

	vector<pair<string, string>> entries;
	for (int i = 0; i < count; ++i) {
		string text;
		cin >> text;
		size_t separator = text.find(':');
		if (!cin || separator == string::npos || separator == 0) {
			return 1;
		}
		entries.push_back({text.substr(0, separator), text.substr(separator + 1)});
	}

	if (entries.empty()) {
		return 0;
	}

	Chaining chaining{};
	Cuckoo cuckoo{};

	const int runs = 10;
	nanoseconds chainInsert{}, chainSearch{}, cuckooInsert{}, cuckooSearch{};
	int chainFound = 0;
	int cuckooFound = 0;

	for (int run = 0; run < runs; ++run) {
		if (run > 0) {
			clearChaining(&chaining);
			clearCuckoo(&cuckoo);
		}

		initChaining(&chaining);
		initCuckoo(&cuckoo);
		chainFound = 0;
		cuckooFound = 0;

		auto start = steady_clock::now();
		for (const auto& entry : entries) {
			putChaining(&chaining, entry.first, entry.second);
		}
		chainInsert += duration_cast<nanoseconds>(steady_clock::now() - start);

		start = steady_clock::now();
		for (const auto& entry : entries) {
			if (getChaining(&chaining, entry.first) != nullptr) {
				++chainFound;
			}
		}
		chainSearch += duration_cast<nanoseconds>(steady_clock::now() - start);

		start = steady_clock::now();
		for (const auto& entry : entries) {
			putCuckoo(&cuckoo, entry.first, entry.second);
		}
		cuckooInsert += duration_cast<nanoseconds>(steady_clock::now() - start);

		start = steady_clock::now();
		for (const auto& entry : entries) {
			if (getCuckoo(&cuckoo, entry.first) != nullptr) {
				++cuckooFound;
			}
		}
		cuckooSearch += duration_cast<nanoseconds>(steady_clock::now() - start);

		if (chainFound != count || cuckooFound != count) {
			cerr << "Ошибка поиска: часть ключей не найдена.\n";
			clearChaining(&chaining);
			clearCuckoo(&cuckoo);
			return 1;
		}
	}

	const double nanosecondsPerMillisecond = 1'000'000.0;
	const double divisor = runs * nanosecondsPerMillisecond;

	cout << fixed << setprecision(6);
	cout << "Число прогонов: " << runs << '\n';
	cout << "Пар в каждом прогоне: " << entries.size() << '\n';
	cout << "Среднее время вставки за прогон, мс: "
		<< chainInsert.count() / divisor << " / " << cuckooInsert.count() / divisor << '\n';
	cout << "Среднее время поиска за прогон, мс: "
		<< chainSearch.count() / divisor << " / " << cuckooSearch.count() / divisor << '\n';
	cout << "Уникальных ключей в таблицах (цепочки / кукушка): "
		<< chaining.size << " / " << cuckoo.size << '\n';

	const string& firstKey = entries.front().first;
	cout << "Найденная пара (цепочки / кукушка): "
		<< firstKey << ':' << *getChaining(&chaining, firstKey) << " / "
		<< firstKey << ':' << *getCuckoo(&cuckoo, firstKey) << '\n';

	// string key;
	// cout << "Ключ для поиска, удаления и повторной вставки: ";
	// if (!(cin >> key)) {
	// 	clearChaining(&chaining);
	// 	clearCuckoo(&cuckoo);
	// 	return 1;
	// }

	// cout << boolalpha;
	// cout << "После вставки (дубли не учитываются):\n";
	// cout << "Цепочки: " << chaining.size << ", кукушка: " << cuckoo.size << '\n';
	// printBloomResults(&chaining, &cuckoo, key);
	// cout << "Поиск " << key << ": " << containsChaining(&chaining, key) << ", " << containsCuckoo(&cuckoo, key) << '\n';
	// cout << "Удаление " << key << ": " << removeChaining(&chaining, key) << ", " << removeCuckoo(&cuckoo, key) << '\n';
	// printBloomResults(&chaining, &cuckoo, key);
	// cout << "Поиск после удаления: " << containsChaining(&chaining, key) << ", " << containsCuckoo(&cuckoo, key) << '\n';

	// addChaining(&chaining, key);
	// addCuckoo(&cuckoo, key);
	// cout << "Вставка " << key << " после удаления: " << containsChaining(&chaining, key) << ", " << containsCuckoo(&cuckoo, key) << '\n';
	// cout << "Вместимость цепочек: " << chaining.capacity << '\n';
	// cout << "Вместимость каждой из двух таблиц кукушки: " << cuckoo.capacity << '\n';

	clearChaining(&chaining);
	clearCuckoo(&cuckoo);
	return 0;
}
