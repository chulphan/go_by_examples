package main

import "fmt"

/*
	for is Go's only looping construct.
	for 는 Go의 유일한 looping construct 이다.
*/

func main() {
	/*
		The most basic type, with a single condition
		하나의 조건과 함께 가장 간단한 타입.
	*/
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	/*
		A classic initial/condition/after for loop.
		가장 전통적인 초기화/조건/후연산 에 대한 loop.
	*/
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	/*
		for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
		조건이 없는 for 문은 loop 의 밖으로 break 하기 전까지 또는 함수가 닫히는 부분으로 부터 return 하지 않으면 계속 반복해서 돌아간다.
	*/
	for {
		fmt.Println("loop")
		break
	}

	/*
		You can also continue to the next iteration of the loop.
		continue 를 사용해서 루프의 다음 반복을 이어나가게 할 수 있다.
	*/
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
