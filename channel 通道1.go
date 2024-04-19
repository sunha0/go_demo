package main

import (
	"fmt"
)

func channel1() {
	var chan1 chan int

	//chan1 = make(chan int)    // 无缓冲区 会阻塞  必须有消费才能往chan写入
	chan1 = make(chan int, 1) // 带缓冲区 1

	chan1 <- 10

	x := <-chan1

	fmt.Println(x)

	_ = len(chan1) //获取chan中元素的数量

	_ = cap(chan1) // 获取chan的容量

	close(chan1)

}

/*
两个gorountine 两个channel
1. 生成0-100 的数字发送到chan1
2. 从chan1 中取出数据计算它的平方 把结果发到ch2中
*/

func channel2() {

	chan1 := make(chan int, 100)
	chan2 := make(chan int, 100)

	go func(chan1 chan int) {
		for i := 0; i <= 100; i++ {

			chan1 <- i
		}
		close(chan1)
	}(chan1)

	go func(chan1 chan int, chan2 chan int) {
		for {
			tmp, ok := <-chan1

			if !ok {
				break
			}
			chan2 <- tmp * tmp

		}
		close(chan2)

	}(chan1, chan2)

	for ret := range chan2 {

		fmt.Println(ret)

	}

}

func main() {

	channel2()

}
