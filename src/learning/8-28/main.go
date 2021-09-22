package main

import (
	"fmt"
	"unicode"
)

func main() {
	test1()
}

func test1() {
	const str = "学习go语言"
	count1, count2 := 0, 0
	for _, s := range str {
		if len(string(s)) >= 3 {
			count1++
		}
		if unicode.Is(unicode.Han, s) {
			count2++
		}
	}
	fmt.Printf("汉字的数量是：%d\n", count1)
	fmt.Printf("汉字的数量是：%d\n", count2)
}

func test2() {
	arr := [...]int{1, 1, 2, 3, 3, 4, 4, 5, 5, 6, 6}
	num := 0
	for _, s := range arr {
		num ^= s
	}
	fmt.Printf("单独的数字是：%d", num)
}

func test3() {
	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Printf("\n")
	}

}

func test4() {
	arr := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, s := range arr {
		sum += s
	}
	fmt.Printf("所有数字的和是：%d", sum)
}

func test5(arr []int, n int) {
	for r, s := range arr {
		for i := r + 1; i < len(arr); i++ {
			if s+arr[i] == n {
				fmt.Printf("(%d, %d)", r, i)
			}
		}

	}
}
