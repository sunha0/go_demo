package main

import (
	"fmt"
	"time"
)

func main() {

	jobs := make(chan int, 100)
	result := make(chan int, 100)

	//分配5个任务
	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	close(jobs)

	// 3个gorountine 去消费job
	for j := 1; j <= 3; j++ {

		go func(j int, jobs chan int, result chan int) {

			for job := range jobs {
				fmt.Printf("work%d 开始做 job%d \n", j, job)
				time.Sleep(time.Millisecond * 60)
				// fmt.Printf("work%d 停止做 job%d \n", j, job)
				result <- job * 2
			}

		}(j, jobs, result)

	}

	for i := 1; i <= 5; i++ {

		res := <-result
		fmt.Println(res)
	}

}
