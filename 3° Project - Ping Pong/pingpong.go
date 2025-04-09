package main

import (
	"fmt"
	"time"
)
func ping(pingCh *chan string, pongCh *chan string){
	for {
		*pingCh <- "ping"
		time.Sleep(1 * time.Second)
		<- *pongCh
	}
}
func pong(pingCh *chan string, pongCh *chan string){
	for {
		msg := <- *pingCh
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
		fmt.Println("pong")
		*pongCh <- "pong"
	}
}
func main()  {

	c1 := make(chan string)
	c2 := make(chan string)

	go ping(&c1, &c2)
	go pong(&c1, &c2)

	var entrada string
	fmt.Scanln(&entrada)
}