package main

import "fmt"

/*
	Go's structs are typed collections of fields.
	They're useful for grouping data together to form records.
	Go 의 struct는 필드들의 타입된 collections이다.
	struct는 레코드 형태로 데이터를 함께 그룹핑하기 유용하다.
*/

/*
	This person struct type has name and age fields.
	이 person struct 타입은 name 과 age 필드를 갖는다.
*/

type person struct {
	name string
	age  int
}

/*
	NewPerson constructs a new person struct with the given name.
	NewPerson 은 새로운 주어진 이름과 함께 새로운 person struct 를 만든다.
*/
func NewPerson(name string) *person {

	/*
		You can safely return a pointer to local variable as a local variable will survive the scope of the function.
		지역 변수가 함수의 범위에서 살아남기 때문에 지역 변수에 안전하게 포인터를 반환할 수 있다.
	*/

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	/*
		이 문법은 struct를 생성한다.
	*/
	fmt.Println(person{"Bob", 20})

	/*
		You can name the fields when initializing a struct.
		struct를 초기화 할 때 필드 이름을 사용할 수 있다.
	*/
	fmt.Println(person{name: "Chul", age: 31})

	/*
		Omitted fields will be zero-valued.
		생략된 필드들은 zero-value 로 될 것이다.
	*/
	fmt.Println(person{name: "ChulChul"})

	/*
		An & prefix yields a pointer to the struct.
		& 접미사는 struct 로의 포인터를 나타낸다.
	*/
	fmt.Println(&person{name: "Ann", age: 40})

	/*
		It's idiomatic to encapsulate new struct creation in constructor functions.
		이 것은 생성자 함수들 내에 새로운 stract 생성을 캡슐화 하는 것의 자연스러운 방법.
	*/
	fmt.Println(NewPerson("ChulKim"))

	/*
		Access struct fields with dot.
		struct 필드에 dot 으로 접근한다.
	*/
	s := person{name: "HELLO~~", age: 45}
	fmt.Println(s.name)

	/*
		You can also use dots with struct pointers - the pointers are automatically dereferenced.
		또 struct 포인터와 함께 점을 사용할 수 있다 - 그 포인터는 자동적으로 역참조 된다.
	*/
	sp := &s
	fmt.Println(sp.age)

	/*
		Structs are mutable.
		Struct는 mutable 이다.(가변)
	*/
	sp.age = 51
	fmt.Println(sp.age)
}
