package main

import "fmt"

/*
	Go supports methods defined on struct types.
	Go는 struct 타입들 상에 정의된 메소드들을 지원한다.
*/
type rect struct {
	width, height int
}

/*
	This area method has a receiver type of *rect.
	이 area 메소드는 *rect 타입의 receiver 를 갖는다.
*/
func (r *rect) area() int {
	return r.width * r.height
}

/*
	Methods can be defined for either pointer or value receiver types.
	Here's an example of a value receiver.
	메소드는 포인터 또는 값 receiver 타입으로 정의 될 수 있다.
	여기 값 receiver 의 예제가 있다
*/
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	/*
		Here we call the 2 methods defined for our struct.
		우리는 우리의 struct를 위해 정의된 2개의 메소드들을 호출한다.
	*/
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	/*
		Go automatically handles conversion between values and pointers for method calls.
		You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
		Go는 메소드 호출에 대해 값과 포인터 변환을 자동적으로 조절한다.
		너는 아마 함수 호출 상에서 복사하는 것을 피하거나 메소드가 받은 struct 를 변경하기 위해서 pointer receiver 를 사용하길 원할수도있다
	*/
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
	fmt.Println("rp: ", rp)
	fmt.Println("*rp: ", *rp)
}
