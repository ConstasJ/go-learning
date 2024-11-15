package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Index   Args")
	for index, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", index, arg)
	}
}
