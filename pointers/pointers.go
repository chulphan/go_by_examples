package main

import "fmt"

/*
	Go supports pointers, allowing you to pass references to values and records within your program.
	Go는 프로그램 내의 레코드들과 값들을 참조할 수 있도록 하는 포인터를 지원한다.
*/

/*
	We’ll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr.
	zeroval has an int parameter, so arguments will be passed to it by value.
	zeroval will get a copy of ival distinct from the one in the calling function.
	우리는 pointer가 zeroval 과 zeroptr 이라는 두 개의 함수가 값들에 어떻게 대조적으로 일하는지 볼거임.
	zeroval 을 int 매개변수를 가져서 인자들은 값에 의해 전달 될거임.
	zeroval은 호출하는 함수 내에 있는 것과 구별되는 ival 의 복사본을 가질것임.
*/

func zeroval(ival int) {
	ival = 0
}

/*
	zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
	The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address.
	Assigning a value to a dereferenced pointer changes the value at the referenced address.
	반대로 zeroptr 은 *int 매개변수를 갖는다. 그 의미는 매개변수가 int 포인터를 취한다는 것임.
	함수 몸통 내에 *iptr 코드는 메모리 주소를 해당 주소의 현재 값으로 역참조한다.
	역참조 된 포인터를 값에 할당하는 것은 참조된 주소에서의 값을 변경한다.
*/
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval", i)
	/*
		The &i syntax gives the memory address of i, i.e. a pointer to i.
		&i 문법은 i의 메모리 주소를 준다. 즉, i로의 포인터임.
	*/
	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	/*
		Pointers can be printed too.
		포인터 역시도 출력할 수 있다.
	*/
	fmt.Println("pointer: ", &i)

	/*
		zeroval doesn’t change the i in main, but zeroptr does because it has a reference to the memory address for that variable.
		zeroval 은 main 내에서 i를 변경하지 않지만 zeroptr은 변수에 대한 메모리 주소를 참조하고 있기 때문에 i를 main 내에서 변경한다.
	*/
}
