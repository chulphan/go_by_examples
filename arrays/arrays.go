package main

import "fmt"

/*
	In Go, an array is a numbered sequence of elements of a specific length.
	Go에서 배열은 명시된 길이의 원소들의 번호가 붙은 수열이다.
*/
func main() {
	/*
		Here we create an array 'a' that will hold exactly 5 ints.
		The type of elements and length are both part of the array's type.
		By default an array is zero-valued, which for ints means 0s.
		여기 우리는 정확히 5개의 정수들을 가지는 배열 'a' 를 만들었음.
		원소들의 타입과 길이는 모두 배열의 타입임!!!!
		자연스럽게 배열은 zero-value임. 즉 ints 의 zero-value 는 0..
	*/
	var a [5]int
	fmt.Println("emp: ", a)

	/*
		We can set a value at an index using the array[index] = value syntax, and get a value with array[index].
		우리는 저 문법을 사용해서 index에 값을 맞출 수 있고 array[index] 라는 것을 통해서 해당 인덱스의 값을 얻을 수 있다.
	*/
	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	/*
		The builtin len returns the length of an array.
		내장된 len 은 배열의 길이를 반환한다.
	*/
	fmt.Println("len: ", len(a))

	/*
		Use this syntax to declare and initialize an array in one line.
		이 문법을 사용하여 배열을 하나의 라인에서 선언하고 초괴화 할 수 있다.
	*/
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	/*
		Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
		배열 타입들은 1차원이지만 타입들을 조합해서 다차원 데이터 구조들을 만들 수 있다.
	*/
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	/*
		Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println
		fmt.Println 으로 배열을 출력하면 [v1 v2 v3 ...] 형태로 나타난다잉~
	*/
}
