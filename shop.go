package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

//сумма заказов
func sum(in map[int][]string, users map[string]int, products map[string]int, orders map[int]string) map[int]int {
	out := make(map[int]int)
	history := make(map[string]int)

	for id, order := range in {
		// ключ - отсортированные лексиграфически продукты
		sort.Strings(order)

		key := strings.Join(order, "")
		//история заказов
		sum, exist := history[key]
		if exist {
			out[id] = sum
		} else {
			new_sum := basicSum(order, products)
			history[key] = new_sum
			out[id] = new_sum
		}

		sell(id, out[id], users, orders)

	}
	return out
}

func basicSum(in []string, products map[string]int) (out int) {
	for _, name := range in {
		price, ok := products[name]
		if ok {
			out += price
		}
	}
	return
}
func addProduct(in map[string]int, products map[string]int) {
	for key, value := range in {
		products[key] = value
	}
}

func updateProduct(in map[string]int, products map[string]int) {
	for name, price := range in {
		_, ok := products[name]
		if ok {
			products[name] = price
		}
	}
}

// Снять деньги у пользователя
func sell(in_order_id int, in_sum int, users map[string]int, orders map[int]string) {
	money, exist := users[orders[in_order_id]]
	if exist && money >= in_sum {
		users[orders[in_order_id]] -= in_sum
	}
}

func sortByKey(mode bool, users map[string]int) (out []string) {
	var keys []string
	for login, _ := range users {
		keys = append(keys, login)
	}
	sortByKey_(keys, mode)
	for _, id := range keys {
		out = append(out, fmt.Sprintf("%v:%v", id, users[id]))
	}
	return out
}

func sortByKey_(in []string, mode bool) []string {
	if mode {
		sort.Strings(in)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(in)))
	}
	return in
}

func sortByValue(mode bool, users map[string]int) (out []string) {
	var (
		keys   []string
		values []int
	)

	for login, money := range users {
		keys = append(keys, login)
		values = append(values, money)
	}

	keys, values = sortByValue_(keys, values, mode)

	for i := 0; i < len(values); i++ {
		out = append(out, fmt.Sprintf("%v:%v", keys[i], values[i]))

	}
	return out
}

func sortByValue_(keys []string, values []int, mode bool) (out_keys []string, out_values []int) {
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

	sortByValue_(keys[:left], values[:left], mode)
	sortByValue_(keys[left+1:], values[left+1:], mode)

	return keys, values
}

//// забыл написать когда сдавал, непонятно как. невнимательность она такая. написал для себя.
//// дописано после дедлайна
//func updateProductName(in map[string]string, products map[string]int) {
//	for name, newName := range in {
//		_, ok := products[name]
//		if ok && len(newName)>0{
//			products[newName] = products[name]
//			delete(products,name)
//		}
//	}
//}
//// дописано после дедлайна
//func addUser(name string,money int, user map[string]int) {
//	_, ok := user[name]
//	if !ok && len(name)>0 {
//		user[name] = money
//	}
//}
//// дописано после дедлайна
//func updateBalance(name string,money int,user map[string]int){
//	_, ok := user[name]
//	if ok && len(name)>0 {
//		user[name] = money
//	}
//}
