package main

import (
	"fmt"
	"time"
)

/*
	A goroutine is a lightweight thread of execution.
	고루틴은 실행의 경량 쓰레드임.
*/

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	/*
		Suppose we have a function call f(s).
		Here's how we'd call that in the usual way, running it synchronously.
		우리가 f(s) 함수를 호출했다고 가정하자.
		우리가 일반적인 방법으로 호출하는 방법이고 동기적으로 실행한다.
	*/
	f("direct")

	/*
		To invoke this function in a goroutine, use go f(s).
		This new goroutine will execute concurrently with the calling one.
		이 함수를 고루틴 내에서 실행시키려면 go f(s) 를 써라
		이 새로운 고루틴은 호출한 것과 동시에 실행될 것이다.
	*/
	go f("goroutine")

	/*
		You can also start a goroutine for an anonymous function call.
		고루틴을 익명 함수 호출에 대해 시작할 수 있음
	*/
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	/*
		Our two function calls are running asynchronously in separate goroutines now.
		Wait for them to finish(for a more robust approach, use a WaitGroup).
		우리의 두 함수 호출은 분리된 고루틴 내에서 비동기적으로 실행된다.
		걔네가 끝날때까지 기다려봐.(더 튼튼한 접근을 위해서 WaitGroup를 사용해라.)
	*/
	time.Sleep(time.Second)
	fmt.Println("done")

	/*
		When we run this program, we see the output of the blocking call first, then the interleaved output of the two goroutines.
		This interleaving reflects the goroutines being run concurrently by the Go runtime.
		이 프로그램을 실행 시킬때, 우리는 블로킹 된 호출의 출력을 처음으로 볼 수 있고 그 담에 두 개의 고루틴의 끼워진 출력을 볼 거임.
		이 끼워진 출력은 고루틴이 고의 런타임에 의해 동시에 실행되고 있음을 반영한다.
	*/
}
