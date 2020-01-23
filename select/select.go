package main

import (
	"fmt"
	"time"
)

/*
	Go’s select lets you wait on multiple channel operations.
	Combining goroutines and channels with select is a powerful feature of Go.
	Go 의 select 는 다중 채널 연산들을 기다리게 한다.
	고루틴과 채널을 select 와 조합하는 것은 Go의 강력한 특징이다.
*/

func main() {
	/*
		For our example we'll select across two channels.
		우리의 예제를 위해서 두 개의 채널들을 선택.
	*/
	c1 := make(chan string)
	c2 := make(chan string)

	/*
		Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.
		각 채널은 어떠한 시간이 흐른 뒤에 동시에 값을 받을 것이다.
		예를 들어, blocking RPC 연산들은 고루틴에서 동시에 실행한다.
	*/
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	/*
		We'll use select to await both of these values simultaneously, printing each one as it arrives.
		우리는 select 를 이들 값 모두 동시에 기다리기 위해 사용하며, 도착하는 것 각각을 출력한다.
	*/
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received: ", msg1)
		case msg2 := <-c2:
			fmt.Println("received: ", msg2)
		}
	}

	/*
		We receive the values "one" and then "two" as expected.
		우리가 기대했던 대로 one을 받고 그 다음에 two를 받는다.
	*/

	/*
		Note that the total execution time is only ~2seconds since both the 1 and 2second Sleeps execute concurrently.
		참고로 총 실행시간은 단 2초이다. 이미 1초와 2초 Sleep 하는건 동시에 실행이 되기 때문이다.
	*/
}
