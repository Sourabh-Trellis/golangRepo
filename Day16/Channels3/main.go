package main

import (
	"fmt"
	"sync"
)

func sender(wg *sync.WaitGroup, ch chan int) {

	for i := 1; i < 6; i++ {
		ch <- i
	}
	wg.Done()
}

func reciever(wg *sync.WaitGroup, ch chan int) {
	// for i := 1; i < 6; i++ {
		
	// 	fmt.Println(<-ch)
	// }
	fmt.Println(<-ch)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go sender(&wg, ch)

	wg.Add(1)
	go reciever(&wg, ch)

	wg.Wait()
}
