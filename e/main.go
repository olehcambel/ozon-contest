package main

// Merge2Channels reads from channels, makes op and writes to out channel.
func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	work := func(x int, task func(int) int) chan int {
		res := make(chan int, 1)
		go func() {
			res <- task(x)
		}()
		return res
	}

	go func() {
		for i := 0; i < n; i++ {
			x1, x2 := <-in1, <-in2
			w1, w2 := work(x1, f), work(x2, f)

			// while waiting for w1, w2 is running in goroutine (non-blocking) and vise versa
			out <- <-w1 + <-w2
		}
	}()
}
