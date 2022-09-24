package main

import "fmt"

func main() {
	fmt.Println(Add(2, 3))
}

func Add(a, b int) int {
	return a + b
}

func Sum(n int) int {
	return n * (1 + n) / 2
}
