#include "core/stack.h"

void processAsteroid(Stack* survivors, int size, char direction) {
	if (direction == 'L') size = -size;

	while (!isEmpty(survivors) && size < 0) {
		int head = stoi(survivors->head->value);

		if (head < 0) break;

		size = head + size;
		pop(survivors);
	}

	if (size != 0) {
		push(survivors, to_string(size));
	}
}

bool readAsteroids(Stack* survivors) {
	int count;
	if (!(cin >> count) || count < 0) return false;

	for (int i = 0; i < count; ++i) {
		int size;
		char direction;

		if (!(cin >> size >> direction)) return false;

		if (size <= 0 || (direction != 'L' && direction != 'R')) return false;

		processAsteroid(survivors, size, direction);
	}

	return true;
}

void printAsteroids(Stack* survivors) {
	Stack reversed;
	init(&reversed);

	while (!isEmpty(survivors)) {
		push(&reversed, survivors->head->value);
		pop(survivors);
	}

	bool first = true;
	cout << "[";

  while (!isEmpty(&reversed)) {
    int value = stoi(reversed.head->value);
    int size = value;
    char direction = 'R';

    if (value < 0) {
      size = -value;
      direction = 'L';
    }

    if (!first) {
      cout << ", ";
    }

    cout << "(" << size << ", " << direction << ")";
    first = false;

    pop(&reversed);
  }

  cout << "]\n";
}

void clearStack(Stack* survivors) {
  while (!isEmpty(survivors)) {
    pop(survivors);
  }
}


int main() {
	Stack survivors;
  init(&survivors);

  if (!readAsteroids(&survivors)) {
    clearStack(&survivors);
    return 1;
  }

  printAsteroids(&survivors);

	return 0;
}