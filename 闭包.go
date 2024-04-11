package main

import (

	//"os"
	//"github.com/gin-gonic/gin"

	"fmt"
)

// 闭包=函数+引用环境

func calc(base int) (func(num1 int) int, func(num2 int) int) {

	a := func(num1 int) int {
		base += num1

		return base
	}

	b := func(num2 int) int {
		base -= num2
		return base
	}

	return a, b

}

func main() {

	add, sub := calc(100)

	a := add(100)
	b := sub(100)

	fmt.Printf("a: %d", a)
	fmt.Printf("b: %d", b)

}
