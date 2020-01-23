package main

import "fmt"

/*
	Go supports recursive functions. Here's a classic factorial example.
	Go 는 재귀함수를 지원한다. 여기 전통적인 팩토리얼 예제가 있다.
*/

/*
	This fact function calls itself until it reaches the base case of fact(0).
	이 fact 함수는 자기 자신을 fact(0) 이 되는 경우가 될 때까지 호출한다.
*/

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func factFor(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	return res
}

func main() {
	fmt.Println(fact(7))
	fmt.Println(factFor(7))
}
