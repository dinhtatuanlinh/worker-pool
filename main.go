package main

import (
	"fmt"
)

func main() {
	number := 10
	numberOfWorker := 5
	jobs := make(chan int, number)
	results := make(chan int, number)
	for i := 0; i < numberOfWorker; i++ {
		go worker(jobs, results)
	}
	for i := 1; i <= number; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < number; j++ {
		fmt.Println(<-results)
	}

	//for i := 1; i <= number; i++ {
	//	fmt.Println(fib(i))
	//}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
