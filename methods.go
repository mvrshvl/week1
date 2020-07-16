package main

import (
	"math/rand"
	"sort"
	"strings"
)

func onePlus(in []int) []int {
	for i := 0; i < len(in); i++ {
		in[i]++
	}
	return in
}
func addFiveLast(in []int) []int {
	in = append(in, 5)
	return in
}

func addFiveFirst(in []int) (out []int) {
	out = append(out, 5)
	out = append(out, in[:]...)
	return
}

func getLast(in []int) (out []int, last int) {
	last = in[len(in)-1]
	out = append(in[0 : len(in)-1])
	return
}

func getFirst(in []int) (out []int, first int) {
	first = in[0]
	out = append(in[1:])
	return
}

func getI(in []int, i_in int) (out []int, i_out int) {
	i_out = in[i_in]
	out = append(in[0:i_in], in[i_in+1:]...)
	return
}

func mergeSlice(in_1 []int, in_2 []int) (out []int) {
	out = append(in_1[:], in_2[:]...)
	return
}

func delete2From1(in_1 []int, in_2 []int) (out []int) {
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

func stepLeft(in []int) []int {
	tmp := in[0]
	for i := 0; i < len(in)-1; i++ {
		in[i] = in[i+1]
	}
	in[len(in)-1] = tmp
	return in
}

func stepLeftI(in []int, count int) []int {
	for j := 0; j < count; j++ {
		tmp := in[0]
		for i := 0; i < len(in)-1; i++ {
			in[i] = in[i+1]
		}
		in[len(in)-1] = tmp
	}
	return in
}

func stepRight(in []int) []int {
	tmp := in[len(in)-1]
	for i := len(in) - 1; i > 0; i-- {
		in[i] = in[i-1]
	}
	in[0] = tmp
	return in
}

func stepRightI(in []int, count int) []int {
	for j := 0; j < count; j++ {
		tmp := in[len(in)-1]
		for i := len(in) - 1; i > 0; i-- {
			in[i] = in[i-1]
		}
		in[0] = tmp
	}
	return in
}

func copySlice(in []int) (out []int) {
	c := make([]int, len(in))
	copy(c, in)
	return c
}

func swapPair(in []int) (out []int) {
	for i := 0; i < len(in)-1; i += 2 {
		tmp := in[i+1]
		in[i+1] = in[i]
		in[i] = tmp
	}
	return in
}

func sortInt(in []int, mode bool) []int {
	if mode {
		sort.Ints(in)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(in)))
	}
	return in
}

func sortString(in []string, mode bool) []string {
	if mode {
		sort.Strings(in)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(in)))
	}
	return in
}

// не понял какая сортировка нужна, поэтому на всякий написал свою

func quickSortInt(in []int, flag bool) []int {
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

	quickSortInt(in[:left], flag)
	quickSortInt(in[left+1:], flag)

	return in
}

func quickSortString(in []string, flag bool) []string {
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

	quickSortString(in[:left], flag)
	quickSortString(in[left+1:], flag)

	return in
}

/*------------------------------------------------------------------*/

func countWords(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Fields(s)
	for _, v := range a {
		m[v]++
	}
	return m
}

func getOnesMaps(in []int) map[int]int {
	m := make(map[int]int)
	for i, _ := range in {
		m[in[i]]++
	}
	return m
}

//тут вообще не понял зачем мапы, может как то задачу по другому понял..
func getTwosMaps(in1, in2 []int) map[int]int {
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
func fibonacci(i int, m map[int]int) int {
	if i == 0 || i == 1 {
		return i
	}
	if m[i] == 0 {
		m[i] = fibonacci(i-1, m) + fibonacci(i-2, m)
	}
	return m[i]
}

//дополнительная
func fibonacci_(i int) map[int]int {
	m := map[int]int{
		1: 1,
	}

	fibonacci(i, m)
	return m
}
