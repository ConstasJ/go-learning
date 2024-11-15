package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, sep := "", ""
	start := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("Time used for string concatenation with for loop:", time.Since(start))
	start = time.Now()
	s = strings.Join(os.Args[1:], " ")
	fmt.Println("Time used for string concatenation with strings.Join:", time.Since(start))
	fmt.Println(s)
}
