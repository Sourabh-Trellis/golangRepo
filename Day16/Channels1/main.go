// Unbuffered channels
package main

import (
	"fmt"
	"sync"
	"time"
)

func one(ch chan string, wg *sync.WaitGroup) {
	fmt.Println("Prepairing to send.")
	// time.Sleep(3 * time.Second)
	ch <- "Value"
	wg.Done()
}

func two(ch chan string, wg *sync.WaitGroup) {
	fmt.Println("Prepairing to recieve.")
	time.Sleep(3 * time.Second)
	val := <-ch
	fmt.Println("Value :", val)
	wg.Done()
}

func main() {
	exChan := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go one(exChan, &wg)

	wg.Add(1)
	go two(exChan, &wg)

	wg.Wait()
}
