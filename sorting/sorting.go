package main

import (
	"fmt"
	"sort"
)

/*
	Go's sort package implements sorting for builtins and user-defined types.
	We'll look at sorting for builtins first.
	Go의 정렬 패키지는 내장된 정렬과 사용자 정의 타입들 정렬에 대해 구현한다.
	우리는 먼저 내장된 것부터 볼 것이다.
*/

func main() {
	/*
		Sort methods are specific to the builtin type; here’s an example for strings.
		Note that sorting is in-place, so it changes the given slice and doesn’t return a new one.
		Sort 메소드는 내장된 타입에만 한정된다; 여기 문자열을 위한 예제가 있다.
		정렬이 제자리에 있으므로 지정된 슬라이스가 변경되고 새 슬라이스가 반환되지 않는다는 점에 유의하자.
	*/
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	/*
		We can also use sort to check if a slice is already in sorted order.
		우리는 또한 slice가 이미 정렬되었는지 확인하기 위해 sort를 사용할 수 있다.
	*/
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
