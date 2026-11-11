#include <iostream>

#include "core/chaining.h"
#include "core/cuckoo.h"

using namespace std;

void printBloomResults(const Chaining* chaining, const Cuckoo* cuckoo, const string& key) {
	cout << "Фильтр цепочек: " << (mightContainBloomFilter(&chaining->filter, key) ? "возможно присутствует" : "точно отсутствует") << '\n';

	cout << "Фильтр кукушки: " << (mightContainBloomFilter(&cuckoo->filter, key) ? "возможно присутствует" : "точно отсутствует") << '\n';
}

int main() {
	int count;
	if (!(cin >> count) || count < 0) {
		return 1;
	}

	Chaining chaining{};
	Cuckoo cuckoo{};
	initChaining(&chaining);
	initCuckoo(&cuckoo);

	for (int i = 0; i < count; ++i) {
		string value;
		if (!(cin >> value)) {
			clearChaining(&chaining);
			clearCuckoo(&cuckoo);
			return 1;
		}
		addChaining(&chaining, value);
		addCuckoo(&cuckoo, value);
	}

	string key;
	cout << "Ключ для поиска, удаления и повторной вставки: ";
	if (!(cin >> key)) {
		clearChaining(&chaining);
		clearCuckoo(&cuckoo);
		return 1;
	}

	cout << boolalpha;
	cout << "После вставки (дубли не учитываются):\n";
	cout << "Цепочки: " << chaining.size << ", кукушка: " << cuckoo.size << '\n';
	printBloomResults(&chaining, &cuckoo, key);
	cout << "Поиск " << key << ": " << containsChaining(&chaining, key) << ", " << containsCuckoo(&cuckoo, key) << '\n';
	cout << "Удаление " << key << ": " << removeChaining(&chaining, key) << ", " << removeCuckoo(&cuckoo, key) << '\n';
	printBloomResults(&chaining, &cuckoo, key);
	cout << "Поиск после удаления: " << containsChaining(&chaining, key) << ", " << containsCuckoo(&cuckoo, key) << '\n';

	addChaining(&chaining, key);
	addCuckoo(&cuckoo, key);
	cout << "Вставка " << key << " после удаления: " << containsChaining(&chaining, key) << ", " << containsCuckoo(&cuckoo, key) << '\n';
	cout << "Вместимость цепочек: " << chaining.capacity << '\n';
	cout << "Вместимость каждой из двух таблиц кукушки: " << cuckoo.capacity << '\n';

	clearChaining(&chaining);
	clearCuckoo(&cuckoo);
	return 0;
}
