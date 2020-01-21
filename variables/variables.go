package main

import "fmt"

func main() {
	/*
		In go, variables are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.
		Go 에서 변수들은 컴파일러에 의해 선언되어지고 사용되어진다.
		예) 함수 호출들의 타입이 맞는지 체크.
	*/

	// var declares 1 or more variables
	// var는 하나 또는 이상의 변수들을 선언한다.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	// 한 번에 여러 개의 변수들을 선언할 수 있다.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	// 고는 초기화 된 변수들의 타입을 추론할 것이다.
	var d = true // boolean 타입으로 추론
	fmt.Println(d)

	/*
		Variables declared without a corresponding initialization are zero-valued.
		For example, the zero value for an int is 0.

		변수들이 초기값과 같이 선언되지 않으면 0값이다.
		예를 들어, int 에 대한 zero value 는 0이다. map = nil
	*/
	var e = true
	fmt.Println(e)

	/*
		The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.

		:= 문법은 변수를 초기화하고 선언하는 것의 간단한 방법이다. 예를 들어 저게 이 경우이다.
	*/
	f := "apple" // <==> var f string = "apple"
	fmt.Println(f)
}
