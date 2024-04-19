package main

import (
	"fmt"
	"sync"
	"time"
)

// demo1 通道误用导致的bug
func demo1() {
	wg := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func(j int) {
			for {
				task,ok := <-ch
				if !ok{
					break
				}
				// 这里假设对接收的数据执行某些操作
				fmt.Println(j, task)
				time.Sleep(time.Millisecond * 60)
			}
			wg.Done()
		}(j)
	}
	wg.Wait()
}

func main() {

	demo1()

}
