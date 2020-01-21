package main

import (
	"fmt"
	"math"
)

/*
	Go supports constats of character, string, boolean, and numeric values.
	Go 는 문자, 문자열, boolean, 수치 값에 대한 상수들을 지원한다.
*/

/*
	const declares a contant value.
	const 는 상수값을 선언한다.
*/
const s string = "constant"

func main() {
	fmt.Println(s)

	/*
		A const statement can appear anywhere a var statement can.
		const 문은 var 문이 나타날 수 있는 어디에서든 나타난다.
	*/
	const n = 5000000000

	/*
		Constant expressions perform arithmetic with arbitrary precision.
		상수 표현식은 임의의 precision과 함께 대수연산을 실행한다.
	*/
	const d = 3e20 / n
	fmt.Println(d)

	/*
		A numeric constant has no type until it's given one, such as by an explicit conversion.
		수치 상수는 명시적 형변환과 같이 타입이 주어지기 전까진 아무 타입을 갖지 않는다.
	*/
	fmt.Println(int64(d))

	/*
		A number can be given a type by using it in a context that requires one,
		such as a variable assignment or function call.
		For example, here math.Sin expects a float64

		숫자는 가변적 할당이나 함수 호출과 같이 하나가 필요한 컨텍스트에서 사용함으로써 타입을 부여할 수 있다. 예를 들어, Math.Sin은 float64 를 기대한다.
	*/
	fmt.Println(math.Sin(n))
}
