package main

import (
	"fmt"
)

/*
	Go offers excellent support for string formatting in the printf tradition.
	Here are some examples of common string fomatting tasks.
	Go 는 문자열 포맷팅을 위한 훌륭한 지원을 printf 전통 내에서 제공한다.
	여기 몇가지 공통적인 문자열 포맷팅 예제들을 보자
*/

type point struct {
	x, y int
}

func main() {
	/*
		Go offers several printing “verbs” designed to format general Go values.
		For example, this prints an instance of our point struct.
		Go 는 일반 Go 값을 포맷하기 위해 설계된 출력 "verbs"를 여러 개 제공한다.
		예를 들어 이 것은 우리의 point struct 의 인스턴스를 출력한다.
	*/
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	/*
		If the value is a struct, the %+v variant will include the struct's field names.
		만약에 값이 struct라면 %+v 는 struct 의 필드 이름들도 포함할 것이다.
	*/
	fmt.Printf("%+v\n", p)

	/*
		The %#v variant prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value.
		%#v 는 값의 Go 문법적 표현을 출력한다. 즉, 값을 제공할 소스 코드 스니펫이다.
	*/
	fmt.Printf("%#v\n", p)

	/*
		To print the type of a value, use %T.
		값의 타입을 출력하기 위해서 %T를 써라.
	*/
	fmt.Printf("%T\n", p)

	/*
		Formatting booleans is straight-forward.
		불린을 포맷팅 하는 것은 간결하다.
	*/
	fmt.Printf("%t\n", true)

	/*
		There are many options for formatting integers.
		Use %d for standard, base-10 formatting.
		정수들을 포맷팅 하기 위한 많은 선택지가 있다.
		기본적으로 %d를 써서 10진수로 포맷팅한다.
	*/
	fmt.Printf("%d\n", 123)

	/*
		This prints a binary representation.
		이 것은 이진 표현을 출력한다.
	*/
	fmt.Printf("%b\n", 14)

	/*
		This prints the character corresponding to the given integer.
		이것은 주어진 정수에 대응하는 문자를 출력한다.
	*/
	fmt.Printf("%c\n", 33)

	/*
		%x provides hex encoding.
		%x 는 16진수 인코딩을 제공한다.
	*/
	fmt.Printf("%x\n", 456)

	/*
		There are also several formatting options for floats.
		For basic decimal formatting use %f.
		또한 float들을 위한 포맷팅 옵션들이 있다.
		기본적인 소수점 포맷팅을 위해서 %f를 써라
	*/
	fmt.Printf("%f\n", 78.9)

	/*
		%e and %E format the float in (slightly different versions of) scientific notation.
		%e 와 %E은 (약간 다른 형태의) 과학적 표기 내에서 float를 포맷한다.
	*/
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	/*
		For basic string printing use %s.
		기본적인 문자열을 출력하기 위해서 %s 를 써라.
	*/
	fmt.Printf("%s\n", "\"string\"")

	/*
		To double-quote strings as in Go source, use %q.
		쌍따옴표 내에 문자열들을 Go 소스처럼 보이기 위해서 %q를 써라.
	*/
	fmt.Printf("%q\n", "\"string\"")

	/*
		As with integers seen earlier, %x renders the string in base-16, with two output characters per byte of input.
		일찍이 봤던 정수들 처럼, %x 는 입력 바이트당 2개의 출력 문자로 base-16의 문자열을 준다.
	*/
	fmt.Printf("%x\n", "hex this")

	/*
		To print a representation of a pointer, use %p.
		포인터의 표현을 출력하려면 %p 를 써라.
	*/
	fmt.Printf("%p\n", &p)

	/*
		When formatting numbers you will often want to control the width and precision of the resulting figure.
		To specify the width of an integer, use a number after the % in the verb.
		By default the result will be right-justified and padded with spaces.
		숫자를 포맷할 때, 너는 종종 결과의 폭과 정밀도(소수점)을 조절하기를 원할 것이다.
		정수의 폭을 지정하려면 % 뒤에 있는 숫자를 사용해라.
		기본적으로 결과는 빈 칸이 채워진 채로 오른쪽으로 정렬 될 것이다.
	*/
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	/*
		You can also specify the width of printed floats, though usually you’ll also want to restrict the decimal precision at the same time with the width.precision syntax.
		출력된 부동소수점의 폭을 명시할 수도 있지만 대게 소수점 자리수도 width.precision 문법과 함께 동시에 제한하고 싶을 수도 있다.
	*/
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	/*
		To left-justify, use the - flag.
		왼쪽 정렬을 위해서 - flag 를 사용해라.
	*/
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
}
