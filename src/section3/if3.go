//IF문(3)
package main

import "fmt"

func main() {
	i := 140

	//if - else if 예제(1)

	if 1 >= 120 {
		fmt.Println("120 이상")
	} else if i >= 100 && i < 120 {
		fmt.Println("100 이상 120 미만")
	} else if i < 100 && i >= 50 {
		fmt.Println("50 이상 100미만")
	} else {
		fmt.Println("50 미만")
	}

}
