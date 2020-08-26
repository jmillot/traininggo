package main

import (
	"fmt"
	"log"
	"time"

	"training.go/imgproc/filter"

	"training.go/imgproc/task"
)

func main() {
	log.Println("Démarrage du programme")
	var f filter.Filter = filter.Grayscale{}
	t := task.NewChanTask("./imgs", "dest", f, 16)
	t.Process()

	start := time.Now()
	t.Process()
	elapsed := time.Since(start)
	fmt.Printf("Image processing took %s\n", elapsed)
}
