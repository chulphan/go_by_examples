package main

import "fmt"

/*
	Maps are Go's built-in associative data type (sometimes called hashes or dicts in other language).
	Map은 Go의 내장된 결합된 데이터 타입이다.
*/

func main() {
	/*
		To create an empty map, use the builtin make: make(map[key-type]val-type)
	*/
	m := make(map[string]int)

	/*
		Set key/value pairs using typical name[key] = val syntax.
		일반적으로 name[key] = val 문법을 사용해서 key/value 쌍을 설정한다.
	*/
	m["k1"] = 7
	m["k2"] = 13

	/*
		Printing a map with e.g. fmt.Println will show all of its key/value pairs.
		map을 출력하는 것. 예를 들어 fmt.Println은 map의 모든 key/value 쌍을 보여줄 것이다.
	*/
	fmt.Println("map: ", m)

	/*
		Get a value for a key with name[key]
		name[key] 로 key 에 대한 값을 얻는다.
	*/
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	/*
		The builtin len returns the number of key/value pairs when called on a map.
		내장 len은 map이 호출 될 때 key/value 순서쌍들의 수를 반환한다.
	*/
	fmt.Println("len: ", len(m))

	/*
		The builtin delete removes key/value pairs from a map.
		내장 delete는 map 으로 부터 key/value 쌍을 제거한다.
	*/
	delete(m, "k2")
	fmt.Println("map: ", m)

	/*
		The optional second return value when getting a value from a map indicates if the key was present in the map.
		This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
		Here we didn't need the value itself, so we ignored it with the blank identifier.
		map 으로 부터 값을 얻을 때 선택적으로 두 번째 반환 값은 map 내에 key 가 지도에 있는지를 나타낸다.
		이것은 key 가 없는 것과 key가 0 또는 "" 와 같은 zero value들을 갖는 것 사이에 차이를 분명히 나타낼 때 사용될 수 있다.
		여기서 우리는 값 자체는 필요가 없으므로 blank 식별자를 이용해서 무시한다.
	*/
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	/*
		You can also declare and initialize a new map in the same line with this syntax.
		이 문법으로 같은 라인 내에서 새로운 map을 선언하고 초기화 할 수 있다.
	*/
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)

	/*
		참고로, fmt.Println을 이용해서 map 을 출력하면 ma[k:v k:v] 형태로 나타난다.
	*/
}
