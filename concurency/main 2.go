package main

import (
	"fmt"
	"os"
	"time"
)

func mail() {
	fmt.Println("In mail")
	go notification()
	file, err := os.Create("mail.txt")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	timestamp := time.Now()
	timeString := timestamp.Format("2006-01-02 15:04:05")
	str := "Hello Sangameshwar,I am pleased to congratulate you on completing Milestone 1 of the Go language Training program.During the training sessions, we covered the following topics so far."
	file.WriteString(str)
	defer file.Close()
	file.WriteString(timeString)
}
func notification() {
	fmt.Println("in Notify")
	file, err := os.Create("notification.txt")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	timestamp := time.Now()
	timeString := timestamp.Format("2006-01-02 15:04:05")
	str := "Training Progress Review:Overall, I am pleased to report that Sangam has made significant progress in his Golang training. He has demonstrated a solid grasp of the foundational concepts and have shown enthusiasm for further learning and development."
	file.WriteString(str)
	file.WriteString(timeString)
	defer file.Close()
}
func main() {
	go mail()
	timestamp := time.Now()
	timeString := timestamp.Format("2006-01-02 15:04:05")
	fmt.Println(timeString)

}
