package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	input io.Reader
)

func inputRef(mode string) {
	switch mode {
	case "offline":
		f, _ := os.Open("input.txt")
		// defer f.Close()
		input = bufio.NewReader(f)
	case "online":
		input = bufio.NewReader(os.Stdin)
	}
}

func main() {
	inputRef("offline")
	var (
		T int
	)

	fmt.Fscan(input, &T)
	for i := 0; i < T; i++ {
		fmt.Println(solver())
	}
}

func solver() int {
	var n int
	fmt.Fscan(input, &n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(input, &s[i])
	}

	fmt.Println(s)

	return 10
}
