package main

import "fmt"

func main() {

	chan1 := make(chan int, 10)

	for i := 0; i <= 10; i++ {

		select {
		case x := <-chan1:
			fmt.Println(x)

		case chan1 <- i:

		default:
			fmt.Println("啥也不干等待")

		}

	}

}
