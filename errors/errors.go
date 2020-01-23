package main

import (
	"errors"
	"fmt"
)

/*

	In Go it’s idiomatic to communicate errors via an explicit, separate return value.
	This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C.
	Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for any other, non-error tasks.
	Go 에서 명시적이고 분리된 반환 값을 통해서 에러와 소통하는 것이 자연스럽다.
	이건 예외가 사용되는 언어인 자바와 루비랑 대조되고 하나의 결과와 값을 쓰는 C와도 대조된다.
	Go의 접근법은 오류를 반환하는 함수를 쉽게 확인하고 다른 오류가 아닌 작업에 사용되는 동일한 언어 구조를 사용하여 이를 처리할 수 있도록 한다.
*/

/*
	By convention, errors are the last return value and have type error, a built-in interface.
	관습적으로 에러들은 마지막으로 값을 반환하고 내장 인터페이스인 error라는 타입을 갖는다.
*/
func f1(arg int) (int, error) {
	/*
		errors.New constructs a basic error value with the given error message.
		errors.New 는 기본적인 에러를 기본적인 에러 값을 주어진 에러 메세지와 만든다.
	*/
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	/*
		A nil value in the error position indicates that there was no error.
		에러 위치 내에 nil 값은 에러가 없다는 것을 나타낸다.
	*/
	return arg + 3, nil
}

/*
	It's possible to use custom types as errors by implementing the Error() method on them.
	Here's a variant on the example above that uses a custom type to explicitly represent an argument error.
	Error() 메소드를 커스텀 타입에 구현하는 것에 의해 커스텀 타입을 에러처럼 사용할 수 있다.
	argument error를 표현하는 다양한 예제들을 보자.
*/
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		/*
			In this case we use &argError syntax to build a new struct,
			supplying values for the two fields arg and prob.
			이 경우에 우리는 &argError 문법을 사용해서 새로운 struct 를 빌드하고
			두 필드 arg와 prob을 위한 값을 제공한다.
		*/
		return -1, &argError{arg, "can't work with it."}
	}
	return arg + 3, nil
}

func main() {
	/*
		The two loops below test out each of our error-returning functions.
		Note that the use of an inline error check on the 'if' line is a common idiom in Go code.
		아래에 두 개의 루프는 우리의 에러 반환 함수를 테스트한다.
		'if' 라인에 인라인 오류 검사를 사용하는 것은 Go 코드의 일반적인 관용어라는 점에 유의하십시오.
	*/
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed...", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed: ", e)
		} else {
			fmt.Println("f2 worked: ", r)
		}
	}

	/*
		If you want to programmatically use the data in a custom error,
		you'll need to get the error as an instance of the custom error type via type assertion.
	*/
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
