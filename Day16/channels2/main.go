package main

import (
	"fmt"
	"sync"
	"time"
)

func sender1(exChan chan string, wg *sync.WaitGroup) {
	fmt.Println("sender 1")
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		exChan <- fmt.Sprintf("%d", i)
	}
	defer wg.Done()
}

func sender2(exChan chan string, wg *sync.WaitGroup) {
	fmt.Println("sender 2")
	for i := 10; i < 20; i++ {
		time.Sleep(1 * time.Second)
		exChan <- fmt.Sprintf("%d", i)
	}
	defer wg.Done()
}

func sender3(exChan chan string, wg *sync.WaitGroup) {
	fmt.Println("sender 3")
	for i := 20; i < 30; i++ {
		time.Sleep(1 * time.Second)
		exChan <- fmt.Sprintf("%d", i)
	}
	defer wg.Done()
}

func sender4(exChan chan string, wg *sync.WaitGroup) {
	fmt.Println("sender 4")
	for i := 30; i < 40; i++ {
		time.Sleep(1 * time.Second)
		exChan <- fmt.Sprintf("%d", i)
	}
	defer wg.Done()
}

func reader1(exChan chan string, wg *sync.WaitGroup) {
	fmt.Println("reader 1")
	for message := range exChan {
		data := fmt.Sprintf("value: %s - buffer size- %d", message, len(exChan))
		fmt.Println("reader 1 - ", data)
	}
	close(exChan)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	exChan := make(chan string, 10)
	fmt.Println("Channel size - ", len(exChan))

	wg.Add(1)
	go reader1(exChan, &wg)

	wg.Add(1)
	go sender1(exChan, &wg)
	wg.Add(1)
	go sender2(exChan, &wg)
	wg.Add(1)
	go sender3(exChan, &wg)
	wg.Add(1)
	go sender4(exChan, &wg)

	wg.Wait()

}
