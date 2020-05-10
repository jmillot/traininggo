package main

import (
<<<<<<< HEAD
=======
	"flag"
>>>>>>> 52174efb162ed6e01dac325a0d73d14f67d95324
	"fmt"
	"os"

	"training.go/Dictionary/dictionary"
)

func main() {
<<<<<<< HEAD
=======
	action := flag.String("action", "list", "Action to perform on the dictionary")

>>>>>>> 52174efb162ed6e01dac325a0d73d14f67d95324
	d, err := dictionary.New("./badger")
	handleError(err)
	defer d.Close()

<<<<<<< HEAD
}

=======
	flag.Parse()
	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	default:
		fmt.Printf("Unknown action: %v\n", *action)
	}
}

func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleError(err)
	fmt.Printf("'%v' added to the dictionary\n", word)
}

func actionDefine(d *dictionary.Dictionary, args []string) {
	word := args[0]
	entry, err := d.Get(word)
	handleError(err)
	fmt.Println(entry)
}

func actionRemove(d *dictionary.Dictionary, args []string) {
	word := args[0]
	err := d.Remove(word)
	handleError(err)
	fmt.Printf("'%v' removed to the dictionary\n", word)
}

func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	handleError(err)
	fmt.Println("Dictionary content")
	for _, word := range words {
		fmt.Println(entries[word])
	}

}
>>>>>>> 52174efb162ed6e01dac325a0d73d14f67d95324
func handleError(err error) {
	if err != nil {
		fmt.Printf("Dictionary error:%v\n", err)
		os.Exit(1)
	}
}
