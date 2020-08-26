package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go ping(c)
	go pong(c)
	go print(c)

	time.Sleep(10 * time.Second)
}

func ping(c chan string) {
	for i := 1; ; i++ {
		c <- fmt.Sprintf("ping %v", i)
	}
}

func pong(c chan string) {
	for i := 100; ; i = +100 {
		c <- fmt.Sprintf("pong %v", i)
	}
}

func print(c chan string) {
	for {
		message := <-c
		fmt.Println(message)
		time.Sleep(1 * time.Second)
	}
}
