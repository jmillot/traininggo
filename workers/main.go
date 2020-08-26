package main

import (
	"fmt"
	"time"
)

func longOperation(duration int) {
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Printf("Operation finished ! Took %d s\n", duration)
}

func main() {
	fmt.Println("Starting first operaton")
	go longOperation(1)

	fmt.Println("Starting second operaton")
	longOperation(2)

	time.Sleep(2 * time.Second)

}
