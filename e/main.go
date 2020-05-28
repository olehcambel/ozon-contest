package main

import (
	"fmt"
	"sync"
	"time"
)

// Merge2Channels reads from channels, makes op and writes to out channel.
func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	for i := 0; i < n; i++ {
		x1, x2 := <-in1, <-in2
		out <- f(x1) + f(x2)
	}
}

func op(x int) int {
	time.Sleep(time.Second * 2)
	return x
}

func main() {
	n := 10
	in1 := make(chan int, n)
	in2 := make(chan int, n)
	out := make(chan int, n)

	for i := 0; i < n; i++ {
		in1 <- i * 10
		in2 <- i + 1
	}

	go Merge2Channels(op, in1, in2, out, n)

	for i := 0; i < n; i++ {
		fmt.Println(<-out)
	}
}

func groups() {
	var wg sync.WaitGroup
	wg.Add(2)

	work := func(id int) {
		defer wg.Done()
		fmt.Printf("Горутина %d начала выполнение \n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Горутина %d завершила выполнение \n", id)
	}

	go work(1)
	go work(2)
	go work(3)

	wg.Wait()
	wg.Wait()

	fmt.Print("done")
}
