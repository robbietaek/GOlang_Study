//Switch 예제 (3)
package main

import "fmt"

func main() {

	//예제1
	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println(a, "는 짝수")
	case 1, 3, 5:
		fmt.Println(a, "는 홀수")
	}

	//예제2 조건에 맞는 문장 아래에 반드시 실행되어야하는 문장이 있다면 fallthrough 를 사용할 수 있다
	switch e := "go"; e {
	case "java":
		fmt.Println("Java!")
		fallthrough
	case "go":
		fmt.Println("go!")
		fallthrough
	case "python":
		fmt.Println("python!")
		fallthrough
	case "javascript":
		fmt.Println("javascript!")
		fallthrough
	case "c#":
		fmt.Println("c#!")
		//fallthrough
	}
}
