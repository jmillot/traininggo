package main

import "fmt"

func main() {
	i := 1
	var p *int = &i

	fmt.Printf("i=%v\n", i)
	fmt.Printf("p=%v\n", p)
}
