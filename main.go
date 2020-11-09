package main

import (
	"fmt"
)

// Channel

// 1. Unbuffered channel
// 2. Buffered channel
// 3. select 
// 4. close channel

func main() {
	/*
	// Unbuffered channel
	ch := make(chan int)

	go func() {
		ch <- 100
		fmt.Println("sent")
	}()

	i := <- ch

	fmt.Println(i)
	fmt.Println("Done")
	*/

	/*
	// Buffered channel
	// Have capacity
	//ch := make(chan int, 0) // <=> Unbuffered Declaration
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3 // Deadlock

	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch) 
	//fmt.Println(<-ch) // Deadlock

	fmt.Println(<-ch) // take value from closed channel
	fmt.Println(<-ch)
	*/

	/*
	// Select is used when our application have many goroutines
	queue := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			queue <- i
		}

		done <- true
	} ()

	// Non-stop for-loop
	for {
		select {
			case v := <- queue:
				fmt.Println(v)
			case <- done:
				fmt.Println("done")
				return
		}
	}*/

	// Demo another close channel example
	// close channel
	queue := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			if i > 5 {
				close(queue)
				break
			} else {
				queue <- i
			}
		}
		//close(queue)
	} ()
	
	for value := range queue {
		fmt.Println(value)
	}

	// never send a value into a closed channel
	q := make(chan int, 10)
	close(q)

	q <- 1
}