//For문(1)
package main

import "fmt"

func main() {
	//반복문 - For
	//Go 에서 유일하게 제공되는 반복문
	//다양한 사용법 숙지

	//예제1

	for i := 0; i < 5; i++ {
		fmt.Println("ex1 :", i)
	}

	/*
		//에러 발생1 (괄호 띄어쓰기 안됨)
		for i:= 0; i < 5; i++
		{

		}
	*/

	/*
		//에러 발생 2 (한줄일때 괄호생략안됨)
		for i:=0; i< 5; i++
			fmt.Println("i :",i)
	*/

	/*

		//예제 2 (무한루프)
		for  {
			fmt.Println("ex2 : Hello Golang!")
		}
	*/

	//예제 3 (Range 용법)
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}

}
