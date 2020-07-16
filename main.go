package main

import "fmt"

func main() {
	/*--------------ДЗ1----------------------------------*/
	//1.1
	//fmt.Println(Sqrt(2))
	//1.2
	//pic.Show(Pic)
	//1.3
	//wc.Test( WordCount)
	//1.4
	//f := fibonacci()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(f())
	//}
	/*---------------------------------------------------*/
	fmt.Println("\n\n1.1\n", "IN", []int{1, 2, 3, 4, 5})
	fmt.Println("OUT", que11([]int{1, 2, 3, 4, 5}))

	fmt.Println("\n\n1.2\n", "IN", []int{1, 2, 3, 4, 5})
	fmt.Println("OUT", que12([]int{1, 2, 3, 4, 5}))

	fmt.Println("\n\n1.3\n", "IN", []int{1, 2, 3, 4, 5})
	fmt.Println("OUT", que13([]int{1, 2, 3, 4, 5}))

	var (
		a []int = []int{1, 2, 3, 4, 5}
		b int
		c []int = []int{4, 5, 9, 10, 1}
	)

	fmt.Println("\n\n1.4\n", "IN", a)
	a, b = que14(a)
	fmt.Println("OUT", a, b)

	a = []int{1, 2, 3, 4, 5}
	fmt.Println("\n\n1.5\n", "IN", a)
	a, b = que15(a)
	fmt.Println("OUT", a, b)

	// fmt.Scan(&i) Нужен ли?
	a = []int{1, 2, 3, 4, 5}
	in_ind := 1
	fmt.Println("\n\n1.6\n", "IN", a, in_ind)
	a, b = que16(a, in_ind)
	fmt.Println("OUT", a, b)

	a = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("\n\n1.7\n", "IN", a, c)
	fmt.Println("OUT", que17(a, c))

	fmt.Println("\n\n1.8\n", "IN", a, c)
	fmt.Println("OUT", que18(a, c))

	fmt.Println("\n\n1.9\n", "IN", a)
	fmt.Println("OUT", que19(a))

	fmt.Println("\n\n1.10\n", "IN", a, in_ind)
	fmt.Println("OUT", que110(a, in_ind))

	fmt.Println("\n\n1.11\n", "IN", a)
	fmt.Println("OUT", que111(a))

	fmt.Println("\n\n1.12\n", "IN", a, in_ind)
	fmt.Println("OUT", que112(a, in_ind))

	fmt.Printf("\n\n1.13\n IN %v ADDR %p\n", a, a)
	que_113 := que113(a)
	fmt.Printf("OUT %v ADDR %p\n", que_113, que_113)

	fmt.Println("\n\n1.14\n", "IN", a)
	fmt.Println("OUT", que114(a))

	fmt.Println("\n\n1.15 INT\n", "IN", a)
	fmt.Println("OUT <", qsort(a, true))
	fmt.Println("OUT >", qsort(a, false))
	fmt.Println("\n\n1.15 STRING\n", "IN", []string{"ba", "aa", "ac", "s", "q"})
	fmt.Println("OUT", lsort([]string{"ba", "aa", "ac", "s", "q"}, true))

	var (
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

	fmt.Println("\n\n2.1\n", "IN", text)
	fmt.Println("OUT", que21(text))

	fmt.Println("\n\n2.2\n", "IN", numbers1)
	fmt.Println("OUT", que22(numbers1))

	fmt.Println("\n\n2.3\n", "IN", numbers1, numbers2)

	fmt.Println("OUT", que23(numbers1, numbers2))
	fmt.Println("OUT", que23(numbers1, numbers2))

	fmt.Println("\n\n2.4\n", "IN", i_fib)
	fmt.Println("OUT", que24_(i_fib))

	fmt.Println("Изначальный список:", que25Get())
	fmt.Println("\n\n2.5 сумма\n", "IN", my_list)
	fmt.Println("OUT", que25Sum(my_list))

	fmt.Println("\n\n2.5 добавление товара\n", "IN", add_list)
	que25Add(add_list)
	fmt.Println("Новый список", que25Get())

	fmt.Println("\n\n2.5 изменение цены\n", "IN", upd_list)
	que25Upd(upd_list)
	fmt.Println("Новый список\n", que25Get())

	fmt.Print("\n\n2.6 супер сумма\n IN\n")

	for id, list := range supersum {
		fmt.Println(id, ":", list)
	}
	fmt.Print("OUT\n")
	for id, sum := range que26(supersum) {
		fmt.Println(id, ":", sum)
	}

	fmt.Print("\n\n2.7 Теперь есть пользователи\n")

	fmt.Println("деньги Васи", getUserMoney("вася"))
	fmt.Println("деньги Пети", getUserMoney("петя"))

	for id, list := range supersum {
		fmt.Println(id, ":", list)
	}
	fmt.Print("OUT\n")
	for id, sum := range que27(supersum) {
		fmt.Println(id, ":", sum, "Заказ оформил:", getUserByOrder(id))
	}
	fmt.Println("деньги Васи", getUserMoney("вася"))
	fmt.Println("деньги Пети", getUserMoney("петя"))

	fmt.Println("\n\n2.8 Теперь есть сортировка")

	fmt.Println("\nЛогин по убыванию")
	que28("k", false)
	fmt.Println("\nЛогин по возрастанию")
	que28("k", true)
	fmt.Println("\nДенюшки по убыванию")
	que28("v", false)
	fmt.Println("\nДенюшки по возрастанию")
	que28("v", true)

}
