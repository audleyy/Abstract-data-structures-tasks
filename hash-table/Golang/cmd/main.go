package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"lab2/hash-table/Golang"
)

func printBloomResults(chaining *hashtable.Set, cuckoo *hashtable.Cuckoo, key string) {
	chainingResult := "точно отсутствует"
	if hashtable.MightContainBloomFilter(&chaining.Filter, key) {
		chainingResult = "возможно присутствует"
	}
	fmt.Println("Фильтр цепочек:", chainingResult)

	cuckooResult := "точно отсутствует"
	if hashtable.MightContainBloomFilter(&cuckoo.Filter, key) {
		cuckooResult = "возможно присутствует"
	}
	fmt.Println("Фильтр кукушки:", cuckooResult)
}

func run() int {
	input := bufio.NewReader(os.Stdin)
	var countText string
	if _, err := fmt.Fscan(input, &countText); err != nil {
		return 1
	}
	count, err := strconv.ParseInt(countText, 10, 32)
	if err != nil || count < 0 {
		return 1
	}

	var chaining hashtable.Set
	var cuckoo hashtable.Cuckoo
	hashtable.InitSet(&chaining)
	hashtable.InitCuckoo(&cuckoo)
	defer hashtable.Clear(&chaining)
	defer hashtable.ClearCuckoo(&cuckoo)

	for i := int64(0); i < count; i++ {
		var value string
		if _, err := fmt.Fscan(input, &value); err != nil {
			return 1
		}
		hashtable.Add(&chaining, value)
		hashtable.AddCuckoo(&cuckoo, value)
	}

	fmt.Print("Ключ для поиска, удаления и повторной вставки: ")
	var key string
	if _, err := fmt.Fscan(input, &key); err != nil {
		return 1
	}

	fmt.Println("После вставки (дубли не учитываются):")
	fmt.Printf("Цепочки: %d, кукушка: %d\n", chaining.Size, cuckoo.Size)
	printBloomResults(&chaining, &cuckoo, key)
	fmt.Printf("Поиск %s: %t, %t\n", key, hashtable.Contains(&chaining, key), hashtable.ContainsCuckoo(&cuckoo, key))
	fmt.Printf("Удаление %s: %t, %t\n", key, hashtable.Remove(&chaining, key), hashtable.RemoveCuckoo(&cuckoo, key))
	printBloomResults(&chaining, &cuckoo, key)
	fmt.Printf("Поиск после удаления: %t, %t\n", hashtable.Contains(&chaining, key), hashtable.ContainsCuckoo(&cuckoo, key))

	hashtable.Add(&chaining, key)
	hashtable.AddCuckoo(&cuckoo, key)
	fmt.Printf("Вставка %s после удаления: %t, %t\n", key, hashtable.Contains(&chaining, key), hashtable.ContainsCuckoo(&cuckoo, key))
	fmt.Printf("Вместимость цепочек: %d\n", chaining.Capacity)
	fmt.Printf("Вместимость каждой из двух таблиц кукушки: %d\n", cuckoo.Capacity)
	return 0
}

func main() {
	os.Exit(run())
}
