package main

import "fmt"

func main() {
	const (
		A = iota * 10
		B
		C
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
