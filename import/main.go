package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("model.json")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 4096)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(data)
	var objMap map[string]interface{}
	err = json.Unmarshal(data, &objMap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(objMap)
}
