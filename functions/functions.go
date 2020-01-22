package main

import "fmt"

/*
	Functions are central in Go. We'll learn about functions with a few different examples.
	함수는 Go에서 가장 중요하다. 몇 가지 예제들을 가지고 함수에 대해 배워보자.
*/

/*
	여기의 함수는 2개의 정수를 가지고 그들의 합을 int로 반환한다.
*/
func plus(a int, b int) int {
	/*
		Go requires explicit returns, i.e. it won't automatically return the value of the last expression.
		Go는 명시적 반환을 요구한다. 즉, 마지막 표현식의 값을 자동적으로 반환하지 않을 것이다.
	*/
	return a + b
}

/*
	When you have multiple consecutive parameters of the same type,
	you may omit the type name for the like-typed parameters up to the final parameters up to the final parameter that declares the type.
	같은 타입의 연속된 parameter들을 여러 개 가지고 있을 때,
	같은 타입을 가지는 마지막 매개변수까지 타입 이름을 생략할 수 있다.
*/
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	/*
		Call a function just as you'd expect, with name(args)
		함수를 호출 하는 방법은 이미 알고 있듯이 name(args) 로 호출한다.
	*/
	res := plus(1, 2)
	fmt.Println("1+2= ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3= ", res)

	/*
		There are several other features to Go functions.
		One is multiple return values, which we'll look at next.
		Go 함수들에게는 몇 가지 다른 특징들이 있다.
		그 중 한가지는 여러 개의 값을 반환할 수 있는데, 다음 장에서 볼거야..
	*/
}
