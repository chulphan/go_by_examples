package main

import "fmt"

/*
	Go supports anonymous functions, which can form closures.
	Anonymous functions are useful when you want to define a function inline without having to name it.
	Go 는 익명함수를 지원하는데, 익명함수는 클로저 형태가 될 수 있다.
	익명함수는 함수의 이름을 지을 필요 없이 함수 내부에 함수를 정의하기를 원할 때에 유용하다.
*/

/*
	This function intSeq returns another function, which we define anonymously in the body of intSeq.
	The returned function closes over the variable i to form a closure.
	이 함수 intSeq는 또 다른 함수를 반환하는데, 그 함수는 우리가 intSeq의 몸체 내에 익명적으로 정의한 것이다.
	반환된 함수는 변수 i를 감싸는 클로저의 형태이다.
*/

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	/*
		We call intSeq, assigning the result (a function) to nextInt.
		This function value captures its own i value, which will be updated each time we call nextInt.
		우리는 intSeq 를 호출해서 함수의 결과를 nextInt 변수에 할당한다.
		이 함수값은 자신의 i값을 점유하고 i값은 nextInt를 호출할 때 마다 갱신될 것이다..
	*/

	nextInt := intSeq()
	/*
		See the effect of the closure by calling nextInt a few times.
		nextInt 를 몇 번 호출하는 것으로 closure의 효과를 볼 수 있다.
	*/
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	/*
		To confirm that the state is unique to that particular function, create and test a new one.
		상태가 특정 함수에 고유하다는 것을 확인하려면 새로운 함수를 만들어서 테스트 해보자.
	*/
	newInts := intSeq()
	fmt.Println(newInts())

	/*
		The last feature of functions we'll look at for now is recursion.
		이제 함수의 마지막 특징인 recursion(재귀)를 볼 것이다.
	*/
}
