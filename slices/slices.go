package main

import "fmt"

/*
	Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.
	Slice는 Go에서 핵심 데이터 타입이다. 수열들에 배열보다 더 강력한 인터페이스를 제공한다.
*/

func main() {
	/*
		Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
		To create an empty slice with non-zero length, use the builtin make.
		Here we make a slice of strings of length 3(initially zero-valued)
		배열과 다르게 slice는 오직 포함하고 있는 원소들에 의해서만 타입이 정해진다.(원소들의 수가 아니라..)
		길이가 0보다 큰 빈 slice 를 만들기 위해서는 내장된 make 를 사용한다.
		여기서 우리는 길이가 3인 문자열들의 slice 를 만든다.
	*/
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	/*
		We can set and get just like with arrays.
		우리는 배열이랑 똑같은 방법으로 slice의 값을 설정하고 얻을 수 있다.
	*/

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	/*
		len returns the length of the slice as expected.
		len 은 기대한 바와 같이 slice의 길이를 반환한다.
	*/
	fmt.Println("len: ", len(s))

	/*
		In addition to these basic operations, slices support several more that make them richer than arrays.
		One is the builtin append, which returns a slice containing one or more new values.
		Note that we need to accept a return value from append as we may get a new slice value.
		slice 는 배열보다 더 풍부하게 만드는 몇 가지 기능을 추가적으로 제공한다.
		그 중 한 가지는 빌트인 append이다. append는 하나 또는 하나 이상의 새로운 값을 포함하는 slice 를 반환한다.
		새로운 slice 값을 얻을 수 있으므로 append 에서 반환 값을 수락해야 함을 참고.
	*/
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	/*
		Slices can also be copy'd. Here we create an empty slice c of the same length as s and copy into c from s.
		slice 또한 copy 되어질 수 있다. 여기서 우리는 s 와 같은 길이인 빈 slice c 를 만들고 s를 c로 복사 할 것이다..
	*/
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	/*
		Slices support a "slice" operator with the syntax slice[low:high].
		For example, this gets a slice of the elements s[2], s[3], and s[4].
		Slice는 slice 연산자와 함께 'slice' 기능을 지원한다.
		예를 들어, 아래 예제는 원소들의 슬라이스인 s[2], s[3], s[4] 를 얻는다.
	*/
	l := s[2:5]
	fmt.Println("sl1: ", l)

	/*
		This slices up to (but excluding) s[5].
		이거는 4까지의 슬라이스를 제공한다.
	*/
	l = s[:5]
	fmt.Println("sl2: ", l)

	/*
		And this slices up from (and including) s[2].
		그리고 이거는 s[2]를 포함해서 끝까지를 slice 한다.
	*/
	l = s[2:]
	fmt.Println("sl3: ", l)

	/*
		We can declare and initialize a variable for slice in a single line as well.
		우리는 한 줄로 slice 에 대한 변수를 선언하고 초기화 할 수 있다.
	*/
	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	/*
		Slices can be composed into multi-dimensional data structures.
		The length of the inner slices can vary, unlike with multi-dimensional arrays.
		slice는 다차원 데이터 구조들로 구성될 수 있다.
		내부 slice들의 길이는 다차원 배열들과는 달리 다양할 수 있다..(당연히 slice의 길이는 동적이 될 수 있으니까)
	*/
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	/*
		Note that while slices are different types than arrays, they are rendered similarly by fmt.Println
		배열들과 slice 들이 다름에도 불구하고 fmt.Println 으로 찍으면 생긴게 비슷하다는걸 참고하셈.
	*/
}
