package main

import (
	"fmt"
	"sync"
	"time"
)

// func main() {

// 	go hello()
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("hi my name is sourabh")

// }

// func hello() {

// 	fmt.Println("hello moto")

// }


func main()  {

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		
		wg.Add(1)
		go task(i+1,&wg)

	}

	wg.Wait()
	fmt.Println("all tasks completed")
	
}
func task(id int,wg *sync.WaitGroup)  {
	
	fmt.Printf("task %d starting......\n",id)
	time.Sleep(time.Second)
	fmt.Printf("task %d ending......\n",id)
	wg.Done() 

}