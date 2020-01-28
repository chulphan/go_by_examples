package main

import (
	"fmt"
	"sort"
)

/*
	Sometimes we’ll want to sort a collection by something other than its natural order.
	For example, suppose we wanted to sort strings by their length instead of alphabetically.
	Here’s an example of custom sorts in Go.
	떄로 우리는 자연적인 순서보다는 어떠한 다른 무언가를 기준으로 collection을 정렬하고 싶을 때가 있을 것이다.
	예를 들어 우리가 문자열을 알파벳 순이 아닌 길이로 정렬하길 원한다고 가정하자.
	여기 Go 내에서 커스텀 정렬을 하는 예제가 있다.
*/

/*
	In order to sort by a custom function in Go, we need a corresponding type.
	Here we’ve created a byLength type that is just an alias for the builtin []string type.
	Go 내에서 커스텀 함수에 의해 정렬을 하기 위해서 대응되는 타입이 필요하다.
	여기 우리는 byLength 타입을 만들었는데 byLength 타입은 단지 내장된 []string type을 위한 alias 이다.
*/

type byLength []string

/*
	We implement sort.Interface - Len, Less, and Swap - on our type so we can use the sort package’s generic Sort function.
	Len and Swap will usually be similar across types and Less will hold the actual custom sorting logic.
	In our case we want to sort in order of increasing string length, so we use len(s[i]) and len(s[j]) here.
	우리는 sort.Interface - Len, Less, 그리고 Swap 을 우리의 타입 상에서 구현해서 sort 패키지의 제네릭 Sort 함수를 사용할 수 있다.
	Len과 Swap 은 대게 타입에 따라 비슷하며 Less는 실제 커스텀 정렬 로직을 가지고 있을 것이다.
	우리의 경우에는 우리는 문자열 길이가 증가하는 순서대로 정렬하길 원하기 때문에 우리는 여기서 len(s[i]) 와 len(s[j]) 를 사용한다.
*/
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	/*
		With all of this in place, we can now implement our custom sort by converting the original fruits slice to byLength, and then use sort.Sort on that typed slice.
		이렇게 모든 것을 갖추게 되면, 우리는 이제 fruits 슬라이스를 byLength 로 변환한 다음, sort.Sort를 타입이 정해진 slice에 사용함으로써 우리의 커스텀 정렬을 구현할 수 있다.
	*/
	fruits := []string{"peach", "banana", "kiwi", "corn1"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	/*
		Running our program shows a list sorted by string length, as desired.
		우리의 프로그램을 실행하면 우리가 원했던 대로 문자열 길이에 의해 정렬된 리스트를 보여준다.
	*/

	/*
		By following this same pattern of creating a custom type, implementing the three Interface methods on that type, and then calling sort.Sort on a collection of that custom type, we can sort Go slices by arbitrary functions.
		사용자 정의 타입을 생성하는 동일한 패턴을 따르고 해당 타입에 대해 세 가지 인터페이스 방법을 구현한 다음 sort.Sort를 호출해라.
		해당 커스텀 타입의 컬렉션을 기준으로 정렬하면 임의의 함수로 Go 슬라이스르 정렬할 수 있다.
	*/
}
