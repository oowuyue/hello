package main

import (
	"fmt"
)

func main() {
	// 声明
	var i int

	// 声明  初始化 推倒
	var j = 2
	h := 2 //函数内

	fmt.Println(i + j + h)

	var middle = 19 / 2
	fmt.Println(middle)

}

func s1(n int) int {
	if n == 1 {
		return 1
	}
	return s1(n-1) + n
}
