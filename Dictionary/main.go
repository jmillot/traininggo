package main

import (
	"fmt"
	"os"

	"training.go/Dictionary/dictionary"
)

func main() {
	d, err := dictionary.New("./badger")
	handleError(err)
	defer d.Close()

}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err)
		os.Exit(1)
	}
}
