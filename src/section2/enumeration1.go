//열거형
package main

import "fmt"

func main() {
	//열거형
	//상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음
	const (
		Jan = 1
		Feb = 2
		Mar = 3
		Apr = 4
		May = 5
		Jun = 6
	)

	fmt.Println("jan : ", Jan)
	fmt.Println("Feb : ", Feb)
	fmt.Println("Mar : ", Mar)
	fmt.Println("Apr : ", Apr)
	fmt.Println("May : ", May)
	fmt.Println("Jun : ", Jun)

}
