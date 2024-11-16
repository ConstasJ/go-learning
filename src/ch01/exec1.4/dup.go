package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	counts := make(map[string]int)
	lineToFile := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineToFile)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineToFile)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t出现文件：%s\n", n, line, strings.Join(lineToFile[line], ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, lineToFile map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if lineToFile[text] == nil {
			lineToFile[text] = []string{f.Name()}
		} else {
			if !slices.Contains(lineToFile[text], f.Name()) {
				lineToFile[text] = append(lineToFile[text], f.Name())
			}
		}
		counts[text]++
	}
}
