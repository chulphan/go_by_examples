package main

import "fmt"

/*
	Go has built-in support for multiple return values.
	This feature is used often in idiomatic Go,
	for example to return both result and error values from a function.
	Go 는 내부적으로 다수의 값들을 반환하는 것을 지원한다.
	이 특징은 Go 내에서 자연스럽게 사용되는데,
	예를 들어 함수로 부터 결과와 에러를 둘 다 반환하는 것이다.
*/

/*
	The (int, int) in this function signature shows that the function returns 2 ints.
*/
func vals() (int, int) {
	return 3, 7
}

func main() {
	/*
		Here we use the 2 different return values from the call with multiple assignment.
		우리는 호출로 부터 반환되는 2 개의 반환 값들을 다중 할당을 통해 사용한다.
	*/
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	/*
		If you only want a subset of the returned values, use the blank identifier _.
		만약 너가 반환 값들 중에 일부만 사용하고 싶으면 blank 식별자 _ 를 사용해라.
	*/
	_, c := vals()
	fmt.Println(c)

	/*
		Accepting a variable number of arguments is another nice feature of Go functions;
		we'll look at this next.
		인자들의 변수 개수를 받아들이는 것은 Go 함수의 또 다른 좋은 특징이다.
		이 다음 장에서 살펴보자.
	*/
}
