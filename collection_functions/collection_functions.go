package main

import (
	"fmt"
	"strings"
)

/*
	We often need our programs to perform operations on collections of data,
	like selecting all items that satisfy a given predicate or mapping all items to a new collection with a custom function.
	우리는 종종 우리의 프로그램들이 주어진 술어를 만족하는 또는 커스텀 함수와 함께 모든 아이템들이 새로운 컬랙션에 매핑되는 모든 아이템들을 선택하는 것과 같은 data의 컬렉션 상에서 연산들을 실행하는 것이 필요하다.
*/

/*
	In some languages it’s idiomatic to use generic data structures and algorithms.
	Go does not support generics; in Go it’s common to provide collection functions if and when they are specifically needed for your program and data types.
	어떤 언어들에서는 제네릭 자료구조와 알고리즘을 사용하는 것이 자연스럽다.
	Go는 제네릭을 지원하지 않는다; Go에서는 프로그램과 데이터 유형에 특별히 필요한 경우 수집 기능을 제공하는 것이 일반적이다.
*/

/*
	Here are some example collection functions for slices of strings.
	You can use these examples to build your own functions.
	Note that in some cases it may be clearest to just inline the collection-manipulating code directly, instead of creating and calling a helper function.
	여기 문자열의 슬라이스를 위한 몇 가지 컬랙션 함수들에 대한 예제가 있다.
	너는 이 예제들을 사용해서 너 자신만의 함수들을 만들 수 있다.
	참고로 이 몇 가지 경우들은 아마 헬퍼 함수를 생성하고 호출하는 것 대신에 컬랙션을 조작하는 코드를 직접 inline 하는 것이 더 명확할 수 있다.
*/

/*
	Index returns the first index of the target string t, or -1 if no match is found.
	Index는 타겟 문자열 t의 첫번째 인덱스를 반환하거나 매치되지 않는다면 -1을 반환한다.
*/
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

/*
	Include returns true if the target string t is in the slice.
	Include는 슬라이스 내에 타겟 문자열 t가 있다면 true를 반환한다.
*/
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

/*
	Any returns true if one of the strings in the slice satisfies the predicate f.
	Any 는 슬라이스 내에 문자열 중 하나가 술부 f를 만족하면 true 를 반환한다.
*/
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

/*
	All returns true if all of the strings in the slice satisfy the predicate f.
	All 은 슬라이스 내에 모든 문자열들이 술부 f를 만족하면 true를 반환한다.
*/
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

/*
	Filter returns a new slice containing all strings in the slice that satisfy the predicate f.
	Filter 슬라이스 내에 술부 f를 만족하는 모든 문자열들을 포함하는 새로운 슬라이스를 반환한다.
*/
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

/*
	Map returns a new slice containing the results of applying the function f to each string in the original slice.
	Map은 원래 슬라이스의 각 문자열에 함수 f를 적용한 결과를 포함하는 새로운 슬라이스를 반환한다.
*/
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	/*
		Here we try out our various collection functions.
		다양한 컬랙션 함수들을 시도해보자~~
	*/
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println(Map(strs, strings.ToUpper))

	/*
		The above examples all used anonymous functions, but you can also use name functions of the correct type.
		위의 예제들은 모두 익명함수들로 쓰여졌지만 정확한 타입의 명명된 함수들을 쓸 수도 있다.
	*/
}
