//Go 특징(1)
package main

import "fmt"

func main() {
	//GO : 모호하거나, 애매한 문법 금지
	//후치(후위) 연산자 허용 (i++), 전치(전위) 연산자 비허용 (++i) -> X
	//증감 연산 반환값 없음 sum += i++ -> X
	//포인터(Pointer 허용, 연산 비허용)
	//주석 (// . /**/) 사용법 숙지

	//예제1
	sum, i := 0, 0

	for i <= 100 {
		//sum += i++
		sum += i
		i++
		//++i 예외발생(전위증감)
	}

	fmt.Println("ex1 : ", sum)

}
