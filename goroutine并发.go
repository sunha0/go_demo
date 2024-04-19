package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(num int) {

	fmt.Printf("hello %d\n", num)
	wg.Done()

}

func main() { // 开启一个主goroutine 执行main函数
	wg.Add(10)
	for i := 0; i < 10; i++ {

		go hello(i) // 开启另一个goroutine 执行hello()
	}
	fmt.Println("hello 执行完了")
	// time.Sleep(2 * time.Second)

	wg.Wait()
}
