package main

import (
	"fmt"
	"time"
)

func main() {

	//get current time
	currentTime := time.Now()
	fmt.Println("Current time:", currentTime)

	//parse time using specific layout
	layout := "2006-01-02 15:04:05"
	timeStamp := "2024-06-23 12:12:12"
	parsedTime, err := time.Parse(layout, timeStamp)
	if err != nil {
		fmt.Println("Time parse Error:", err)
		return
	}
	fmt.Println("Parsed time :", parsedTime)

	//add and subtract duration from time
	oneHour, _ := time.ParseDuration("1h")
	oneDay, _ := time.ParseDuration("24h")

	futureTime := currentTime.Add(oneHour)
	fmt.Println("Future time :", futureTime)

	pastTime := currentTime.Add(-oneDay)
	fmt.Println("Past Time :", pastTime)

	//calculate difference between two times
	duration := parsedTime.Sub(currentTime)
	fmt.Println("Time difference :", duration)

}
