package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortByName(t *testing.T) {
	// нет такого
}

func TestHashOrder(t *testing.T) {
	//встроен в calculate
}

func TestCalculateOrderWithCache(t *testing.T) {

	users = map[string]int{
		"вася": 300,
	}

	res := sum(map[int][]string{}, users, nil, nil)

	t.Log(res)

	if users["вася"] != 300 {
		t.Fatal()
	}

	res = sum(map[int][]string{}, users, nil, map[int]string{})

	if users["вася"] != 300 {
		t.Fatal()
	}

	res = sum(map[int][]string{1: {"a"}, 2: {"b"}}, users, map[string]int{"a": 1, "b": 10}, map[int]string{1: "вася"})

	if users["вася"] != 299 {
		t.Fatal()
	}

	users["вася"] = 300

	res = sum(map[int][]string{1: {"a"}, 2: {"b"}}, users, map[string]int{"a": 1, "b": 10}, map[int]string{1: "вася", 2: "вася"})

	if users["вася"] != 289 {
		t.Fatal()
	}

	users["вася"] = 300

	res = sum(map[int][]string{1: {"xxx"}}, users, map[string]int{"a": 1, "b": 10}, map[int]string{1: "вася"})

	if users["вася"] != 300 {
		t.Fatal()
	}

	users["вася"] = 300

	res = sum(map[int][]string{1: {"a"}, 2: {"xxx"}, 3: {"b"}}, users, map[string]int{"a": 1, "b": 10}, map[int]string{1: "вася", 2: "вася", 3: "вася"})

	if users["вася"] != 289 {
		t.Fatal()
	}
	users["вася"] = 300

	res = sum(map[int][]string{1: {"aa"}}, users, map[string]int{"a": 1, "b": 10}, map[int]string{1: "вася"})

	if users["вася"] != 300 {
		t.Fatal()
	}
}

func TestAddItem(t *testing.T) {

	new := map[string]int{}
	addProduct(map[string]int{}, map[string]int{})

	if !reflect.DeepEqual(new, map[string]int{}) {
		t.Fatal()
	}

	new = map[string]int{"": 1}
	products = map[string]int{}
	addProduct(new, products)

	if !reflect.DeepEqual(new, products) {
		t.Fatal()
	}

	new = map[string]int{"a": 5}
	products = map[string]int{
		"a": 1,
		"b": 10,
	}
	addProduct(new, products)

	if !reflect.DeepEqual(products, map[string]int{
		"a": 1,
		"b": 10,
	}) {
		t.Errorf(fmt.Sprint(products))
	}

	new = map[string]int{"xxx": 5}
	products = map[string]int{
		"a": 1,
		"b": 10,
	}
	addProduct(new, products)

	if !reflect.DeepEqual(products, map[string]int{
		"a":   1,
		"b":   10,
		"xxx": 5,
	}) {
		t.Fatal()
	}
}

//как то забыл написать 3 простых метода, исправленное лежит в другой ветке
//func TestChangePrice(t *testing.T) {
//
//	new := map[string]int{}
//	updateProduct(map[string]int{}, map[string]int{})
//
//	if !reflect.DeepEqual(new, map[string]int{}) {
//		t.Fatal()
//	}
//
//	new = map[string]int{"":1}
//	products = map[string]int{}
//	updateProduct(new, products)
//
//	if !reflect.DeepEqual(products, map[string]int{}) {
//		t.Fatal()
//	}
//
//	new = map[string]int{"a":5}
//	products = map[string]int{
//		"a":     1,
//		"b":    10,
//	}
//	updateProduct(new, products)
//
//	if !reflect.DeepEqual(products, map[string]int{
//		"a":     5,
//		"b":    10,
//	}) {
//		t.Fatal(products)
//	}
//
//	new = map[string]int{"xxx":5}
//	products = map[string]int{
//		"a":     1,
//		"b":    10,
//	}
//	updateProduct(new, products)
//
//	if !reflect.DeepEqual(products, map[string]int{
//		"a":     1,
//		"b":    10,
//	}) {
//		t.Fatal()
//	}
//}
//
//func TestChangeName(t *testing.T) {
//
//	new := map[string]string{}
//	shop:=map[string]int{}
//	updateProductName(new, shop)
//
//	if !reflect.DeepEqual(shop, map[string]int{}) {
//		t.Fatal()
//	}
//	shop= map[string]int{"a":10}
//	new = map[string]string{ "a":""}
//	updateProductName(new,shop)
//
//	if !reflect.DeepEqual(shop, map[string]int{"a":10}) {
//		t.Fatal()
//	}
//
//	new = map[string]string{"a":"aa"}
//	shop = map[string]int{"a":1,"b":10}
//	updateProductName(new,shop )
//
//	if !reflect.DeepEqual(shop, map[string]int{"aa":1,"b":10}) {
//		t.Fatal()
//	}
//
//	new = map[string]string{"xxx":"a"}
//	shop =  map[string]int{"a":1,"b":10}
//		updateProductName(new,shop)
//
//	if !reflect.DeepEqual(shop, map[string]int{"a":1,"b":10}) {
//		t.Fatal()
//	}
//}
//
//func TestAddAccount(t *testing.T) {
//
//	users = map[string]int{
//	}
//
//	addUser("",10,users)
//
//	if !reflect.DeepEqual(users, map[string]int{}) {
//		t.Fatal()
//	}
//
//	users = map[string]int{
//		"a" : 1,
//		"b" : 10,
//	}
//
//	addUser("a",10,users)
//	if !reflect.DeepEqual(users, map[string]int{"a":1,"b":10}) {
//		t.Fatal()
//	}
//
//	addUser("xxx",100,users)
//	if !reflect.DeepEqual(users, map[string]int{"a":1,"b":10,"xxx":100}) {
//		t.Fatal()
//	}
//
//}
//
//func TestChangeBalance(t *testing.T) {
//
//	users = map[string]int{
//	}
//
//	updateBalance("",10,users)
//
//	if !reflect.DeepEqual(users, map[string]int{}) {
//		t.Fatal()
//	}
//
//	users = map[string]int{
//		"a" : 1,
//		"b" : 10,
//	}
//
//	updateBalance("a",10,users)
//	if !reflect.DeepEqual(users, map[string]int{"a":10,"b":10}) {
//		t.Fatal()
//	}
//
//	users = map[string]int{
//		"a" : 1,
//		"b" : 10,
//	}
//
//	updateBalance("a",1,users)
//	if !reflect.DeepEqual(users, map[string]int{"a":1,"b":10}) {
//		t.Fatal()
//	}
//
//	updateBalance("xxx",100,users)
//	if !reflect.DeepEqual(users, map[string]int{"a":1,"b":10}) {
//		t.Fatal()
//	}
//}

func TestSortAccounts(t *testing.T) {
	testCases := []struct {
		name             string
		accounts         map[string]int
		sortBy           bool
		expectedAccounts []string
	}{
		// name
		{
			"SortByNameAsc. empty Accounts",
			map[string]int{},
			true,
			[]string{},
		},
		{
			"SortByNameAsc. single Account ",
			map[string]int{
				"a": 10,
			},
			true,
			[]string{"a:10"},
		},
		{
			"SortByNameAsc. already sorted",
			map[string]int{
				"a":     1,
				"b":     10,
				"d":     12,
				"d10":   11,
				"xxx_1": 22,
			},
			true,
			[]string{
				"a:1",
				"b:10",
				"d:12",
				"d10:11",
				"xxx_1:22",
			},
		},
		{
			"SortByNameAsc. already sorted in reverse order",
			map[string]int{
				"xxx_1": 22,
				"d10":   11,
				"d":     12,
				"b":     10,
				"a":     1,
			},
			true,
			[]string{
				"a:1",
				"b:10",
				"d:12",
				"d10:11",
				"xxx_1:22",
			},
		},
		{
			"SortByNameAsc. random order",
			map[string]int{
				"d10":   11,
				"a":     1,
				"xxx_1": 22,
				"b":     10,
				"d":     12,
			},
			true,
			[]string{
				"a:1",
				"b:10",
				"d:12",
				"d10:11",
				"xxx_1:22",
			},
		},

		{
			"SortByNameDesc. empty Accounts",
			map[string]int{},
			false,
			[]string{},
		},
		{
			"SortByNameDesc. single Account ",
			map[string]int{
				"a": 10,
			},
			false,
			[]string{"a:10"},
		},
		{
			"SortByNameDesc. already sorted",
			map[string]int{
				"xxx_1": 22,
				"d10":   11,
				"d":     12,
				"b":     10,
				"a":     1,
			},
			false,
			[]string{
				"xxx_1:22",
				"d10:11",
				"d:12",
				"b:10",
				"a:1",
			},
		},
		{
			"SortByNameDesc. already sorted in reverse order",
			map[string]int{
				"a":     1,
				"b":     10,
				"d":     12,
				"d10":   11,
				"xxx_1": 22,
			},
			false,
			[]string{
				"xxx_1:22",
				"d10:11",
				"d:12",
				"b:10",
				"a:1",
			},
		},
		{
			"SortByNameDesc. random order",
			map[string]int{
				"d10":   11,
				"a":     1,
				"xxx_1": 22,
				"b":     10,
				"d":     12,
			},
			false,
			[]string{
				"xxx_1:22",
				"d10:11",
				"d:12",
				"b:10",
				"a:1",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := sortByKey(tc.sortBy, tc.accounts)

			if !reflect.DeepEqual(res, tc.expectedAccounts) {
				t.Fatalf("got\t\t%v\nwant\t%v", res, tc.expectedAccounts)
			}
		})
	}

}
func TestSortAccountsBalance(t *testing.T) {

	testCases := []struct {
		name             string
		accounts         map[string]int
		sortBy           bool
		expectedAccounts []string
	}{
		// balance

		{
			"SortByBalanceAsc. empty Accounts",
			map[string]int{},
			true,
			[]string{},
		},
		{
			"SortByBalanceAsc. single Account",
			map[string]int{
				"a": 10,
			},
			true,
			[]string{"a:10"},
		},
		{
			"SortByBalanceAsc. already sorted",
			map[string]int{
				"a":     1,
				"b":     10,
				"d10":   11,
				"d":     12,
				"xxx_1": 22,
			},
			true,
			[]string{
				"a:1",
				"b:10",
				"d10:11",
				"d:12",
				"xxx_1:22",
			},
		},
		{
			"SortByBalanceAsc. already sorted with duplicates",
			map[string]int{
				"a":     1,
				"b":     10,
				"d10":   11,
				"d11":   11,
				"d":     12,
				"xxx_1": 22,
				"xxx_2": 22,
			},
			true,
			[]string{
				"a:1",
				"b:10",
				"d10:11",
				"d11:11",
				"d:12",
				"xxx_1:22",
				"xxx_2:22",
			},
		},
		{
			"SortByBalanceAsc. already sorted in reverse order",
			map[string]int{
				"xxx_1": 22,
				"d":     12,
				"d10":   11,
				"b":     10,
				"a":     1,
			},
			true,
			[]string{
				"a:1",
				"b:10",
				"d10:11",
				"d:12",
				"xxx_1:22",
			},
		},
		{
			"SortByBalanceAsc. already sorted in reverse order with duplicated",
			map[string]int{
				"xxx_2": 22,
				"xxx_1": 22,
				"d":     12,
				"d11":   11,
				"d10":   11,
				"b":     10,
				"a":     1,
			},
			true,
			[]string{
				"a:1",
				"b:10",
				"d10:11",
				"d11:11",
				"d:12",
				"xxx_1:22",
				"xxx_2:22",
			},
		},
		{
			"SortByBalanceAsc. random order",
			map[string]int{

				"d10":   11,
				"a":     1,
				"xxx_1": 22,
				"b":     10,
				"d":     12,
			},
			true,
			[]string{
				"a:1",
				"b:10",
				"d10:11",
				"d:12",
				"xxx_1:22",
			},
		},
		{
			"SortByBalanceAsc. random order with duplicates",
			map[string]int{
				"d10":   11,
				"d11":   11,
				"a":     1,
				"xxx_1": 22,
				"xxx_2": 22,
				"b":     10,
				"d":     12,
			},
			true,
			[]string{
				"a:1",
				"b:10",
				"d10:11",
				"d11:11",
				"d:12",
				"xxx_1:22",
				"xxx_2:22",
			},
		},

		{
			"SortByBalanceDesc. empty Accounts",
			map[string]int{},
			false,
			[]string{},
		},
		{
			"SortByBalanceDesc. single Account",
			map[string]int{
				"a": 10,
			},
			false,
			[]string{"a:10"},
		},
		{
			"SortByBalanceDesc. already sorted",
			map[string]int{
				"xxx_1": 22,
				"d":     12,
				"d10":   11,
				"b":     10,
				"a":     1,
			},
			false,
			[]string{
				"xxx_1:22",
				"d:12",
				"d10:11",
				"b:10",
				"a:1",
			},
		},
		{
			"SortByBalanceDesc. already sorted with duplicates",
			map[string]int{
				"xxx_1": 22,
				"xxx_2": 22,
				"d":     12,
				"d11":   11,
				"d10":   11,
				"b":     10,
				"a":     1,
			},
			false,
			[]string{
				"xxx_1:22",
				"xxx_2:22",
				"d:12",
				"d10:11",
				"d11:11",
				"b:10",
				"a:1",
			},
		},
		{
			"SortByBalanceDesc. already sorted in reverse order",
			map[string]int{
				"a":     1,
				"b":     10,
				"d10":   11,
				"d":     12,
				"xxx_1": 22,
			},
			false,
			[]string{
				"xxx_1:22",
				"d:12",
				"d10:11",
				"b:10",
				"a:1",
			},
		},
		{
			"SortByBalanceDesc. already sorted in reverse order with duplicated",
			map[string]int{

				"a":     1,
				"b":     10,
				"d10":   11,
				"d11":   11,
				"d":     12,
				"xxx_1": 22,
				"xxx_2": 22,
			},
			false,
			[]string{
				"xxx_2:22",
				"xxx_1:22",
				"d:12",
				"d10:11",
				"d11:11",
				"b:10",
				"a:1",
			},
		},
		{
			"SortByBalanceDesc. random order",
			map[string]int{
				"d10":   11,
				"a":     1,
				"xxx_1": 22,
				"b":     10,
				"d":     12,
			},
			false,
			[]string{
				"xxx_1:22",
				"d:12",
				"d10:11",
				"b:10",
				"a:1",
			},
		},
		{
			"SortByBalanceDesc. random order with duplicates",
			map[string]int{
				"d10":   11,
				"d11":   11,
				"a":     1,
				"xxx_1": 22,
				"xxx_2": 22,
				"b":     10,
				"d":     12,
			},
			false,
			[]string{
				"xxx_2:22",
				"xxx_1:22",
				"d:12",
				"d10:11",
				"d11:11",
				"b:10",
				"a:1",
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := make([]string, 0, 0)

			res = sortByValue(tc.sortBy, tc.accounts)

			if !reflect.DeepEqual(res, tc.expectedAccounts) {
				t.Fatalf("got\t\t%v\nwant\t%v", res, tc.expectedAccounts)
			}
		})
	}
}
