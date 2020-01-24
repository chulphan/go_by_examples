package main

import (
	"fmt"
	"time"
)

/*
	Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time.
	Implementing timeouts in Go is easy and elegant thanks to channels and select.
	Timeouts은 외부 자원에 연결하거나 실행 시간을 제한해야 하는 프로그램에 중요하다.
	Go에서 Timeout을 구현하는 것은 채널과 select 덕분에 쉽고 우아하다.
*/

func main() {
	/*
		For our example, suppose we're executing an external call that returns its result on a channel c1 after 2s.
		Note that the channel is buffered, so the send in the goroutine is nonblocking.
		This is a common pattern to prevent goroutine leaks in case the channel is never read.
		우리의 예제를 위해서 2초 후 채널 c1에서 결과를 반환하는 외부 호출을 실행한다고 가정해보자.
		참고로 channel 은 buffered 라서 고루틴 내에 발신은 논블로킹이다.
		이것은 채널을 읽지 않을 경우에 고루틴의 낭비를 방지하는 보통의 패턴이다.
	*/
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	/*
		Here's the select implementing a timeout.
		res := <-c1 awaits the result and <-Time.After awaits a value to be sent after the timeout of 1s.
		Since select proceeds with the first receive that's ready, we'll take the timeout case if the operation takes more than the allowed 1s.
		여기 timeout을 구현하는 select이다.
		res := <-c1 은 결과를 기다리고 <-Time.After 는 1초의 타임아웃 후에 보내질 값을 기다린다.
		이미 select 는 준비된 첫번째 수신을 계속할 것이고, 만약에 허락된 1초 이상이 걸리는 연산이라면 timeout 케이스를 취할 것이다.
	*/
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	/*
		If we allow a longer timeout of 3s, then the receive from c2 will succeed and we'll print the result.
		만약 3초 이상이 걸리도록 허락한다면, c2 로 부터 수신이 성공될 것이고 그 결과를 출력할 것이다.
	*/
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

	/*
		Running this program shows the first operation timing out and the second succeeding.
		이 프로그램을 실행하는 것은 첫번째는 timing out을 연산하는 두번째는 성공한 것을 보여준다.
	*/
}
