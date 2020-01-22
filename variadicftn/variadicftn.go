package main

import "fmt"

/*
	Variadic functions can be called with any number of trailing arguments.
	For example, fmt.Println is a common variadic function.
	Variadic 함수는 아무 개수의 인자들을 이용하여 호출될 수 있다.
	예를 들어 fmt.Println 은 공통적인 variadict 함수이다.
*/

/*
	Here's a function that will take an arbitrary number of ints as arguments.
	여기 함수는 임의의 정수들을 인자로서 취한다.
*/
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	/*
		Variadic functions can be called in the usual way with individual arguments.
		Variadic 함수들은 개별 인자들과 함께 보통과 같은 방법으로 호출된다.
	*/
	sum(1, 2)
	sum(1, 2, 3)

	/*
		If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
		만약 이미 slice 내에 다수의 인자들을 가지고 있다면 func(slice...) 를 사용해서 variadic 함수에 그것들을 적용할 수 있다.
	*/
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	/*
		Another key aspect of functions in Go is their ability to form closures, which we'll look at next.
		Go에서 또 다른 함수들의 핵심 측면은 그들의 형태가 closure 가 될 수 있다는 것이다. 다음 장에서 살펴볼거임^^
	*/
}
