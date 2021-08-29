package main

import (
	"fmt"
	"sort"
)

func main() {
	test4()
}

func test1() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4]
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}

func test2() {
	arr := [...]int{8, 2, 6, 4, 3, 7}
	a := arr[:]
	sort.Ints(a)
	fmt.Println(a)
	fmt.Println(arr)
}

func test3() {
	s := "how do you do"
	m := make(map[string]int, 5)
	tmp := ""
	for _, a := range s {
		if string(a) == " " {
			m[tmp]++
			tmp = ""
		} else {
			tmp += string(a)
		}
	}
	m[tmp]++
	for u, v := range m {
		fmt.Printf("%v=%d\n", u, v)
	}
}

func test4() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
