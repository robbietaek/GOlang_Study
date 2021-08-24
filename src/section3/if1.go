//IF문(1)
package main

import "fmt"

func main() {
	//제어문(조건문)
	//IF 문 : 반드시 Boolean 검사 -> 1,0 (사용 불가 : 자동 형 변환 불가)
	//소괄호 사용 X

	var a int = 20
	b := 20

	// 예제1
	if a >= 15 {
		fmt.Println("15 이상이다")
	}

	//예제2
	if b >= 25 {
		fmt.Println("25 이상이다")
	}

	/*
		//에러 발생1 (괄호를 띄어쓰면안됨)
		if b >= 25
		{

		}
	*/

	//에러 발생2 (괄호 생략 불가)
	/*
		if b >= 25
			fmt.Println("25이상")
	*/

	//에러 발생3 (true false 가 아님)
	/*
		if c:= 1; c{
			fmt.Println("true")
		}
	*/

	/*
		//에러 발생4 (c 는 1회성 변수이기때문에 외부에서 사용불가능)
		if c:= 40; c > 35{
			fmt.Println("35이상")
		}
		c += 20

	*/

}
