package main

import (
	"fmt"
	"time"
)

/*
	Switch statements express conditionals across many branches.
	Switch 문은 많은 분기들을 가로지르는 조건들을 표현한다.
*/

func main() {
	/*
		Here's a basic switch
		기본적인 switch 문.
	*/
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	/*
		You can use commas to separate multiple expressions in the same case statement.
		We use the optional default case in this example as well.
		콤마를 사용해서 같은 case 문 내에 여러 표현식들을 분리할 수 있다.
		이 예제에서는 default case 을 사용할 것임.(optional)
	*/
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	/*
		switch without an expression is an alternate way to express if/else logic.
		Here we also show how the case expressions can be non-constants.
		표현식 없는 switch 문은 if/else 로직을 표현하기 위해 관련된 방법.
		여기서 우리는 또 case 표현문들이 어떻게 상수가 안될 수 있는지도 보일 것임.
	*/
	t := time.Now()
	switch {
	// switch-case 를 if 문 쓰는 것처럼 사용 가능. 어차피 case 는 상수밖에 안되는거임
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	/*
		A type switch compares types instead of values. You can use this to discover the type of an interface value.
		In this example, the variable t will have the type corresponding to its clause.
		switch 타입은 값들 대신에 타입들을 비교한다. 이것을 인터페이스 값의 타입을 발견하는 데에 사용할 수 있다.
		이 예제에서는 변수 t는 이들 절에 대응하는 타입을 가질 것이다.
	*/
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey~~")
}
