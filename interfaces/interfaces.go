package main

import (
	"fmt"
	"math"
)

/*
	Interfaces are named collections of method signatures.
	Interface는 메소드 서명의 명명된 컬렉션이다. (method signature: 메소드 선언 부분)
*/

/*
	Here's a basic interface for geometric shapes.
	기하학의 형체에 대한 기본적인 인터페이스.
*/
type geometry interface {
	area() float64
	perim() float64
}

/*
	For our example we'll implement this interface on rect and circle types.
	우리의 예제는 인터페이스에 있는 rect 와 circle 타입을 구현할 것이다.
*/
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

/*
	To implement an interface in Go, we just need to implement all the methods in the interface.
	Here we implement geometry on rects.
	Go 에서 인터페이스를 구현하기 위해서 우린 그저 interface 내에 모든 메소드들을 구현하는 것만 필요.
	rects 상에서의 geometry 구현이다.
*/
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

/*
	The implementation for circles.
	circle 에 대한 구현
*/
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

/*
	If a variable has an interface type, then we can call methods that are in the named interface.
	Here's a generic measure function taking advantage of this to work on any geometry.
	만약 변수가 인터페이스 타입을 가지면, 우리는 명명된 인터페이스 내에 메소드들을 호출할 수 있다.
	제네릭 measure 함수는 어떠한 geometry 상에서 동작한다는 이점을 갖는다.
*/
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	/*
		The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
		circle과 rect struct 타입들은 둘 다 geometry 인터페이스를 구현한다. 그래서 우리는 이들 struct들의 인스턴스를 measure의 인자로써 사용할 수 있다.
	*/
	measure(r)
	measure(c)
}
