package main

import (
	"fmt"
	"time"
)

func main() {

	presenttime := time.Now()

	fmt.Println(presenttime)

	fmt.Println(presenttime.Format(" Monday 15:04:05 01-02-2006"))

	cratedDate := time.Date(2020,time.June,23,12,30,40,0,time.UTC)

	fmt.Println(cratedDate)
	fmt.Println(cratedDate.Format("01-02-2006 15:04:05 Monday"))

}
