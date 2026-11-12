.PHONY: task1 task2 task3 task4 task5 task6 task7

task1:
	g++ -std=c++17 stack-based/asteroids.cpp stack-based/core/stack.cpp -o task1

task2:
	g++ -std=c++17 set/main.cpp set/core/set.cpp \
		hash-table/core/chaining.cpp hash-table/core/bloom.cpp \
		hash-functions/hash.cpp array-based/core/array.cpp -o task2

task3:
	g++ -std=c++17 set-partition/partition.cpp set/core/set.cpp \
		hash-table/core/chaining.cpp hash-table/core/bloom.cpp \
		hash-functions/hash.cpp array-based/core/array.cpp -o task3

task4:
	g++ -std=c++17 array-based/subarray.cpp array-based/core/array.cpp -o task4

task5:
	g++ -std=c++17 huffman/huffman.cpp huffman/core/tree.cpp huffman/core/forwardList.cpp -o task5

task6:
	g++ -std=c++17 hash-table/main.cpp hash-table/core/chaining.cpp \
		hash-table/core/cuckoo.cpp hash-table/core/bloom.cpp \
		hash-functions/hash.cpp array-based/core/array.cpp -o task6

task7:
	g++ -std=c++17 lfu-cache/main.cpp -o task7
