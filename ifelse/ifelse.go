package main

import "fmt"

/*
	Branching with if and else in Go is straight-forward.
	Go 에서 if와 else로 분기를 나누는 것은 간단하다.
*/

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	/*
		You can have an if statement without an else.
		if 문에 else 가 없어도 된다.
	*/
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	/*
		A statement can precede conditionals; any variables declared in this statement are available in all branches.
		하나의 문장은 조건에 앞설 수 있다; 이 문장 내에서 선언된 모든 변수들은 모든 분기들에서 유효하다.
	*/
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	/*
		Note that you don't need parentheses around conditions in Go, but that the braces are required.
		Go 에서는 조건문에 () 를 칠 필요가 없다. 하지만 {} 는 필요하다.
	*/
}
