package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func que11(in []int) []int {
	for i := 0; i < len(in); i++ {
		in[i]++
	}
	return in
}
func que12(in []int) []int {
	in = append(in, 5)
	return in
}

func que13(in []int) (out []int) {
	out = append(out, 5)
	out = append(out, in[:]...)
	return
}

func que14(in []int) (out []int, last int) {
	last = in[len(in)-1]
	out = append(in[0 : len(in)-1])
	return
}

func que15(in []int) (out []int, first int) {
	first = in[0]
	out = append(in[1:])
	return
}

func que16(in []int, i_in int) (out []int, i_out int) {
	i_out = in[i_in]
	out = append(in[0:i_in], in[i_in+1:]...)
	return
}

func que17(in_1 []int, in_2 []int) (out []int) {
	out = append(in_1[:], in_2[:]...)
	return
}

func que18(in_1 []int, in_2 []int) (out []int) {
	for i := 0; i < len(in_2); i++ {
		for j := 0; j < len(in_1); j++ {
			if in_2[i] != in_1[j] {
				out = append(out, in_1[j])
			}
		}
		in_1 = []int{}
		in_1 = append(in_1, out[:]...)
		out = []int{}
	}
	return in_1
}

func que19(in []int) []int {
	tmp := in[0]
	for i := 0; i < len(in)-1; i++ {
		in[i] = in[i+1]
	}
	in[len(in)-1] = tmp
	return in
}

func que110(in []int, count int) []int {
	for j := 0; j < count; j++ {
		tmp := in[0]
		for i := 0; i < len(in)-1; i++ {
			in[i] = in[i+1]
		}
		in[len(in)-1] = tmp
	}
	return in
}

func que111(in []int) []int {
	tmp := in[len(in)-1]
	for i := len(in) - 1; i > 0; i-- {
		in[i] = in[i-1]
	}
	in[0] = tmp
	return in
}

func que112(in []int, count int) []int {
	for j := 0; j < count; j++ {
		tmp := in[len(in)-1]
		for i := len(in) - 1; i > 0; i-- {
			in[i] = in[i-1]
		}
		in[0] = tmp
	}
	return in
}

func que113(in []int) (out []int) {
	c := make([]int, len(in))
	copy(c, in)
	return c
}

func que114(in []int) (out []int) {
	for i := 0; i < len(in)-1; i += 2 {
		tmp := in[i+1]
		in[i+1] = in[i]
		in[i] = tmp
	}
	return in
}

// не понял какая сортировка нужна, поэтому на всякий написал свою
func que115Int(in []int, mode bool) []int {
	if mode {
		sort.Ints(in)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(in)))
	}
	return in
}

func que115String(in []string, mode bool) []string {
	if mode {
		sort.Strings(in)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(in)))
	}
	return in
}

func qsort(in []int, flag bool) []int {
	if len(in) < 2 {
		return in
	}

	left, right := 0, len(in)-1
	pvt := rand.Int() % len(in)

	in[pvt], in[right] = in[right], in[pvt]

	for i, _ := range in {
		if flag {
			if in[i] < in[right] {
				in[left], in[i] = in[i], in[left]
				left++
			}
		} else {
			if in[i] > in[right] {
				in[left], in[i] = in[i], in[left]
				left++
			}
		}
	}

	in[left], in[right] = in[right], in[left]

	qsort(in[:left], flag)
	qsort(in[left+1:], flag)

	return in
}

func lsort(in []string, flag bool) []string {
	if len(in) < 2 {
		return in
	}

	left, right := 0, len(in)-1
	pvt := rand.Int() % len(in)

	in[pvt], in[right] = in[right], in[pvt]

	for i, _ := range in {
		if flag {
			if in[i] < in[right] {
				in[left], in[i] = in[i], in[left]
				left++
			}
		} else {
			if in[i] > in[right] {
				in[left], in[i] = in[i], in[left]
				left++
			}
		}
	}

	in[left], in[right] = in[right], in[left]

	lsort(in[:left], flag)
	lsort(in[left+1:], flag)

	return in
}

/*------------------------------------------------------------------*/

func que21(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Fields(s)
	for _, v := range a {
		m[v]++
	}
	return m
}

func que22(in []int) map[int]int {
	m := make(map[int]int)
	for i, _ := range in {
		m[in[i]]++
	}
	return m
}

//тут вообще не понял зачем мапы, может как то задачу по другому понял..
func que23(in1, in2 []int) map[int]int {
	m := make(map[int]int)
	for elem1 := range in1 {
		flag := false
		for elem2 := range in2 {
			if in1[elem1] == in2[elem2] {
				flag = true
				break
			}
		}
		if flag == true {
			m[in1[elem1]]++
		}
	}
	return m
}

//
//func que23_(in1,in2 map[int]int)(out []int){
//	for key1, _ := range in1 {
//		publ ,ok := in2[key1]
//		if ok {
//			out = append(out,publ)
//		}
//	}
//	return out
//}

//основная
func que24(i int, m map[int]int) int {
	if i == 0 || i == 1 {
		return i
	}
	if m[i] == 0 {
		m[i] = que24(i-1, m) + que24(i-2, m)
	}
	return m[i]
}
//дополнительная
func que24_(i int) map[int]int {
	m := map[int]int{
		1: 1,
	}

	que24(i, m)
	return m
}

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
)

func que25Get() map[string]int {
	return products
}
func que25Sum(in []string) (out int) {
	for _, name := range in {
		price, ok := products[name]
		if ok {
			out += price
		}
	}
	return
}

func que25Add(in map[string]int) {
	for key, value := range in {
		products[key] = value
	}
}

func que25Upd(in map[string]int) {
	for name, price := range in {
		_, ok := products[name]
		if ok {
			products[name] = price
		}
	}
}

func que26(in map[int][]string) map[int]int {
	out := make(map[int]int)
	history := make(map[string]int)

	for id, order := range in {
		// ключ - отсортированные лексиграфически продукты
		order = lsort(order, true)
		key := strings.Join(order, "")
		//история заказов
		sum, exist := history[key]
		if exist {
			out[id] = sum
		} else {
			new_sum := que25Sum(order)
			history[key] = new_sum
			out[id] = new_sum
		}
	}
	return out
}

func que27(in map[int][]string) map[int]int {
	out := make(map[int]int)
	history := make(map[string]int)

	for id, order := range in {
		// ключ - отсортированные лексиграфически продукты
		order = lsort(order, true)
		key := strings.Join(order, "")
		//история заказов
		sum, exist := history[key]
		if exist {
			out[id] = sum
		} else {
			new_sum := que25Sum(order)
			history[key] = new_sum
			out[id] = new_sum
		}

		que27Sell(id, out[id])

	}
	return out
}

func que27Sell(in_order_id int, in_sum int) {
	money, exist := users[orders[in_order_id]]
	if exist && money >= in_sum {
		users[orders[in_order_id]] -= in_sum
	}
}

func getUserMoney(in string) int {
	money, exist := users[in]
	if exist {
		return money
	} else {
		return 0
	}
}

func getUserByOrder(in int) string {
	return orders[in]
}

//mode тип сортировки
//field ключ или значение
func que28(field string, mode bool) {
	var keys []string
	var values []int
	for login, money := range users {
		keys = append(keys, login)
		values = append(values, money)
	}
	if field == "k" {
		que115String(keys, mode)
		for _, id := range keys {
			fmt.Println(id, users[id])
		}
	} else if field == "v" {
		keys, values = sortByValue(keys, values, mode)
		for i := 0; i < len(values); i++ {
			fmt.Println(keys[i], values[i])
		}
	}
}

func sortByValue(keys []string, values []int, mode bool) (out_keys []string, out_values []int) {
	if len(values) < 2 {
		return keys, values
	}

	left, right := 0, len(values)-1
	pvt := rand.Int() % len(values)

	values[pvt], values[right] = values[right], values[pvt]
	keys[pvt], keys[right] = keys[right], keys[pvt]

	for i, _ := range values {
		if mode {
			if values[i] < values[right] {
				values[left], values[i] = values[i], values[left]
				keys[left], keys[i] = keys[i], keys[left]

				left++
			}
		} else {
			if values[i] > values[right] {
				values[left], values[i] = values[i], values[left]
				keys[left], keys[i] = keys[i], keys[left]

				left++
			}
		}
	}

	values[left], values[right] = values[right], values[left]
	keys[left], keys[right] = keys[right], keys[left]

	sortByValue(keys[:left], values[:left], mode)
	sortByValue(keys[left+1:], values[left+1:], mode)

	return keys, values
}
