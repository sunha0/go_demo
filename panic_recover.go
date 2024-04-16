package main

import "fmt"

func panic_recover() {

	defer func() {

		err := recover()

		if err != nil {

			fmt.Printf("捕获异常err: %v", err)
		}
	}()

	panic("异常出错了")
}

func main() {
	panic_recover()
}
