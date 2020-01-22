package main

import "fmt"

/*
	range iterates over elements in a variety of data structures.
	Let's see how to use range with some of the data structures we've already learned.
	range 는 다양한 데이터 구조들 내에 elements 를 순회한다.
	우리가 이미 배운 데이터 구조를 통해서 range 를 어떻게 사용하는지 보장
*/

func main() {
	/*
		Here we use range to sum the numbers in slice. Arrays work like this too.
		우리는 range 를 사용해서 slice 내에 숫자들을 더할 것이다.
		배열도 이와 똑같이 작동한다.
	*/
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	/*
		range on arrays and slices provides both the index and value for each entry.
		Above we didn't need the index, so we ignored it with the blank identifier _.
		Sometimes we actually want the indexes though.
		배열과 slice 상의 range는 각 위치에 해당하는 index 와 값을 제공한다.
		위의 예제에서는 index 가 필요하지 않아서 blank 식별자인 _ 을 이용하여 무시했다.
		가끔은 우리가 index 들을 원할수도 있다.
	*/
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	/*
		range on map iterates over key/value pairs.
		map 상에 range 는 key/value 쌍을 순회한다.
	*/
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	/*
		range can also iterate over just the keys of a map.
		range 는 또한 map 의 key들만 순회할 수도 있다.
	*/
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	/*
		range on strings iterates over Unicode code points.
		The first value is the starting byte index of the rune and the second the rune itself.
		문자열 상에 range 는 Unicode code point들을 순회한다.
		첫번째 값은 rune의 시작하는 byte 인덱스이고 두번째는 rune 그 자체이다.
	*/
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
