package main

import "fmt"

func main() {
	const (
		_ = iota
		A
		B
		C
	)

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		_
		PLATINUM
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	fmt.Println(DEFAULT)
	fmt.Println(SILVER)
	fmt.Println(PLATINUM)

}
