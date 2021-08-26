package main

import "fmt"

func main() {
	//예제1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	//예제2
	sum2, i := 0, 0

	for i <= 100 {
		sum2 += i
		i++
		//j := i++ //Go에서 후치연산은 반환값 X
	}
	fmt.Println(sum2)

	//예제 3
	sum3, i := 0, 0

	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println(sum3)

	//예제4
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4 : ", i, j)
	}

	/*
		//에러 발생
		for i,j := 0,0 ;i<= 10; i++, j += 10{
			fmt.Println("ex4 : ",i,j)
		}
	*/

}
