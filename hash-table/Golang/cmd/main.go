package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

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

	type entry struct {
		key   string
		value string
	}
	var entries []entry
	for i := int64(0); i < count; i++ {
		var pair string
		_, err := fmt.Fscan(input, &pair)
		key, value, found := strings.Cut(pair, ":")
		if err != nil || !found || key == "" {
			fmt.Fprintln(os.Stderr, "Ожидается пара в формате ключ:значение.")
			return 1
		}
		entries = append(entries, entry{key: key, value: value})
	}

	var chaining hashtable.Set
	var cuckoo hashtable.Cuckoo
	hashtable.InitSet(&chaining)
	hashtable.InitCuckoo(&cuckoo)
	defer hashtable.Clear(&chaining)
	defer hashtable.ClearCuckoo(&cuckoo)

	var chainInsert, chainSearch, cuckooInsert, cuckooSearch int64
	chainFound := 0
	cuckooFound := 0

	if len(entries) > 0 {
		start := time.Now()
		for _, pair := range entries {
			hashtable.Put(&chaining, pair.key, pair.value)
		}
		chainInsert = time.Since(start).Nanoseconds()

		start = time.Now()
		for _, pair := range entries {
			if _, found := hashtable.Get(&chaining, pair.key); found {
				chainFound++
			}
		}
		chainSearch = time.Since(start).Nanoseconds()

		start = time.Now()
		for _, pair := range entries {
			hashtable.PutCuckoo(&cuckoo, pair.key, pair.value)
		}
		cuckooInsert = time.Since(start).Nanoseconds()

		start = time.Now()
		for _, pair := range entries {
			if _, found := hashtable.GetCuckoo(&cuckoo, pair.key); found {
				cuckooFound++
			}
		}
		cuckooSearch = time.Since(start).Nanoseconds()
	}

	if chainFound != len(entries) || cuckooFound != len(entries) {
		fmt.Fprintln(os.Stderr, "Ошибка поиска: часть ключей не найдена.")
		return 1
	}
	fmt.Println("Эмпирическое сравнение (цепочки / кукушка, обе таблицы с фильтром Блума)")
	fmt.Printf("Пар на метод: %d\n", len(entries))
	fmt.Printf("Вставка с учётом перестроений, нс: %d / %d\n", chainInsert, cuckooInsert)
	fmt.Printf("Поиск вставленных ключей, нс: %d / %d\n", chainSearch, cuckooSearch)
	fmt.Printf("Уникальных ключей в таблицах (цепочки / кукушка): %d / %d\n", chaining.Size, cuckoo.Size)

	if len(entries) > 0 {
		firstKey := entries[0].key
		chainValue, _ := hashtable.Get(&chaining, firstKey)
		cuckooValue, _ := hashtable.GetCuckoo(&cuckoo, firstKey)
		fmt.Printf("Найденная пара (цепочки / кукушка): %s:%s / %s:%s\n",
			firstKey, chainValue, firstKey, cuckooValue)
	}

	// fmt.Print("Ключ для поиска, удаления и повторной вставки: ")
	// var key string
	// if _, err := fmt.Fscan(input, &key); err != nil {
	// 	return 1
	// }

	// fmt.Println("После вставки (дубли не учитываются):")
	// fmt.Printf("Цепочки: %d, кукушка: %d\n", chaining.Size, cuckoo.Size)
	// printBloomResults(&chaining, &cuckoo, key)
	// fmt.Printf("Поиск %s: %t, %t\n", key, hashtable.Contains(&chaining, key), hashtable.ContainsCuckoo(&cuckoo, key))
	// fmt.Printf("Удаление %s: %t, %t\n", key, hashtable.Remove(&chaining, key), hashtable.RemoveCuckoo(&cuckoo, key))
	// printBloomResults(&chaining, &cuckoo, key)
	// fmt.Printf("Поиск после удаления: %t, %t\n", hashtable.Contains(&chaining, key), hashtable.ContainsCuckoo(&cuckoo, key))

	// hashtable.Add(&chaining, key)
	// hashtable.AddCuckoo(&cuckoo, key)
	// fmt.Printf("Вставка %s после удаления: %t, %t\n", key, hashtable.Contains(&chaining, key), hashtable.ContainsCuckoo(&cuckoo, key))
	// fmt.Printf("Вместимость цепочек: %d\n", chaining.Capacity)
	// fmt.Printf("Вместимость каждой из двух таблиц кукушки: %d\n", cuckoo.Capacity)
	return 0
}

func main() {
	os.Exit(run())
}
