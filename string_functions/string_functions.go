package main

/*
	The standard library’s strings package provides many useful string-related functions.
	Here are some examples to give you a sense of the package.
	기본 라이브러리의 strings 패키지는 문자열과 관련된 많은 함수들을 제공한다.
	패키지에 대한 감각을 줄 수 있는 몇 가지 예제들을 보자.
*/

import (
	"fmt"
	s "strings"
)

/*
	We alias fmt.Println to a shorter name as we'll use it a lot below.
	fmt.Println 을 짧은 이름으로 alias 하여서 아래에서 쓰기 편하게끔 한다.
*/
var p = fmt.Println

func main() {
	/*
		Here’s a sample of the functions available in strings.
		Since these are functions from the package, not methods on the string object itself, we need pass the string in question as the first argument to the function.
		You can find more functions in the strings package docs.
		문자열 내에서 사용 가능한 함수들의 샘플이다.
		이미 이 함수들은 문자열 객체 자체의 메소드가 아니라 패키지로 부터 온 함수들이기 때문에 우리는 문자열을 함수의 첫번째 인자로 전달해야한다.
		너는 strings 패키지 문서 내에서 더 많은 함수들을 찾을 수 있다.
	*/
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index: ", s.Index("test", "e"))
	p("Join: ", s.Join([]string{"8579", "5014"}, "-"))
	p("Repeat: ", s.Repeat("a", 5))
	p("Replace: ", s.Replace("foo", "o", "0", -1))
	p("Replace: ", s.Replace("foo", "o", "0", 1))
	p("Split: ", s.Split("a-b-c-d-e", "-"))
	p("ToLower: ", s.ToLower("TEST"))
	p("ToUpper: ", s.ToUpper("test"))

	p()

	/*
		Not part of strings, but worth mentioning here, are the mechanisms for getting the length of a string in bytes and getting a byte by index.
		문자열들의 일부가 아니라 여기서 언급할 가치가 있는 것은 문자열의 길이를 byte로 얻고 index별 byte를 얻는 매카니즘이다.
	*/
	p("Len: ", len("hello"))
	p("Char: ", "hello"[1])

	/*
		Note that 'len' and indexing above work at the byte level. Go uses 'UTF-8' encoded strings, so this is often useful as-is.
		If you’re working with potentially multi-byte characters you’ll want to use encoding-aware operations.
		byte 수준에서 위처럼 len 과 인덱싱을 하는것에 유의해라. Go 는 UTF-8로 인코딩 된 문자열을 사용하기 때문에 이건 종종 유용하다.
		잠재적으로 multi-byte 문자들을 사용하는 경우 인코딩 인식 작업을 사용할 수 있다.
	*/
}
