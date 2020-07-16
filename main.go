package main

import (
	"fmt"
)

var (
	products = map[string]int{
		"apple":     20,
		"orange":    30,
		"banana":    40,
		"pineapple": 50,
	}
	users = map[string]int{
		"вася": 300,
		"петя": 30000000,
		"анна": 5000,
	}

	orders = map[int]string{
		1: "вася",
		2: "петя",
		3: "вася",
	}

	a []int = []int{1, 2, 3, 4, 5}
	b int
	c []int = []int{4, 5, 9, 10, 1}

	text = "ONE " +
		"TWO TWO " +
		"THREE THREE THREE " +
		"FOUR FOUR FOUR FOUR"
	numbers1 = []int{1, 6, 3, 4, 5, 6}
	numbers2 = []int{1, 2, 3, 4, 7, 8, 9, 10, 11}
	i_fib    = 10
	my_list  = []string{"apple", "orange"}
	add_list = map[string]int{
		"pear":   60,
		"carrot": 80,
	}
	upd_list = map[string]int{
		"banana":    100,
		"pineapple": 100,
	}

	supersum = map[int][]string{
		1: {"banana", "orange"},
		2: {"orange", "banana"},
		3: {"pineapple", "apple", "pineapple"},
	}
)

func main() {
	fmt.Println("\n\n1.1\n", "IN", []int{1, 2, 3, 4, 5})
	fmt.Println("OUT", onePlus([]int{1, 2, 3, 4, 5}))

	fmt.Println("\n\n1.2\n", "IN", []int{1, 2, 3, 4, 5})
	fmt.Println("OUT", addFiveLast([]int{1, 2, 3, 4, 5}))

	fmt.Println("\n\n1.3\n", "IN", []int{1, 2, 3, 4, 5})
	fmt.Println("OUT", addFiveFirst([]int{1, 2, 3, 4, 5}))

	fmt.Println("\n\n1.4\n", "IN", a)
	a, b = getLast(a)
	fmt.Println("OUT", a, b)

	a = []int{1, 2, 3, 4, 5}
	fmt.Println("\n\n1.5\n", "IN", a)
	a, b = getFirst(a)
	fmt.Println("OUT", a, b)

	// fmt.Scan(&i) Нужен ли?
	a = []int{1, 2, 3, 4, 5}
	in_ind := 1
	fmt.Println("\n\n1.6\n", "IN", a, in_ind)
	a, b = getI(a, in_ind)
	fmt.Println("OUT", a, b)

	a = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("\n\n1.7\n", "IN", a, c)
	fmt.Println("OUT", mergeSlice(a, c))

	fmt.Println("\n\n1.8\n", "IN", a, c)
	fmt.Println("OUT", delete2From1(a, c))

	fmt.Println("\n\n1.9\n", "IN", a)
	fmt.Println("OUT", stepLeft(a))

	fmt.Println("\n\n1.10\n", "IN", a, in_ind)
	fmt.Println("OUT", stepLeftI(a, in_ind))

	fmt.Println("\n\n1.11\n", "IN", a)
	fmt.Println("OUT", stepRight(a))

	fmt.Println("\n\n1.12\n", "IN", a, in_ind)
	fmt.Println("OUT", stepRightI(a, in_ind))

	fmt.Printf("\n\n1.13\n IN %v ADDR %p\n", a, a)
	que_113 := copySlice(a)
	fmt.Printf("OUT %v ADDR %p\n", que_113, que_113)

	fmt.Println("\n\n1.14\n", "IN", a)
	fmt.Println("OUT", swapPair(a))

	fmt.Println("\n\n1.15 INT\n", "IN", a)
	fmt.Println("OUT <", quickSortInt(a, true))
	fmt.Println("OUT >", quickSortInt(a, false))
	fmt.Println("\n\n1.15 STRING\n", "IN", []string{"ba", "aa", "ac", "s", "q"})
	fmt.Println("OUT", quickSortString([]string{"ba", "aa", "ac", "s", "q"}, true))

	fmt.Println("\n\n2.1\n", "IN", text)
	fmt.Println("OUT", countWords(text))

	fmt.Println("\n\n2.2\n", "IN", numbers1)
	fmt.Println("OUT", getOnesMaps(numbers1))

	fmt.Println("\n\n2.3\n", "IN", numbers1, numbers2)

	fmt.Println("OUT", getTwosMaps(numbers1, numbers2))

	fmt.Println("\n\n2.4\n", "IN", i_fib)
	fmt.Println("OUT", fibonacci_(i_fib))

	fmt.Println("Изначальный список:", products)
	fmt.Println("\n\n2.5 сумма\n", "IN", my_list)
	fmt.Println("OUT", basicSum(my_list, products))

	fmt.Println("\n\n2.5 добавление товара\n", "IN", add_list)
	addProduct(add_list, products)
	fmt.Println("Новый список", products)

	fmt.Println("\n\n2.5 изменение цены\n", "IN", upd_list)
	updateProduct(upd_list, products)
	fmt.Println("Новый список\n", products)

	fmt.Print("\n\n2.6,2.7 супер сумма и пользователи\n IN\n")

	for id, list := range supersum {
		fmt.Println(id, ":", list)
	}

	fmt.Println("деньги Васи", users["вася"])
	fmt.Println("деньги Пети", users["петя"])

	fmt.Print("OUT\n")
	for id, sum := range sum(supersum, users, products, orders) {
		fmt.Println(id, ":", sum, "Заказ оформил:", orders[id])
	}
	fmt.Println("деньги Васи", users["вася"])
	fmt.Println("деньги Пети", users["петя"])

	fmt.Println("\n\n2.8 сортировка")

	fmt.Println("\nЛогин по убыванию")
	sortByKey(false, users)
	fmt.Println("\nЛогин по возрастанию")
	sortByKey(true, users)
	fmt.Println("\nДенюшки по убыванию")
	sortByValue(false, users)
	fmt.Println("\nДенюшки по возрастанию")
	sortByValue(true, users)
}
