package main

import "fmt"

/*
	When using channels as function parameters, you can specify if a channel is meant to only send or receive values.
	This specificity increases the type-safety of the program.
	채널들을 함수의 파라미터로 사용할 때에는 채널이 오직 값을 발신만 하는지 또는 수신만 하는지를 명시할 수 있다.
	이 명시는 프로그램의 type-safety 를 증가시킨다.
*/

/*
	This ping function only accepts a channel for sending values.
	It would be a compile-time error to try to receive on this channel.
	이 ping 함수는 오직 값을 보내는 것에 대한 체널만 받는다.
	이렇게 하면 이 채널이 값을 받으려고 시도할 때에 컴파일 타임 오류를 낼 것이다.
*/

func ping(pings chan<- string, msg string) {
	pings <- msg
}

/*
	The pong function accepts one channel for receives(pings) and a second for sends(pongs).
	pong 함수는 수신을 위한 하나의 채널을 받고(pings) 두번째는 보내는 것에 대한 채널을 받는다.(pongs)
*/

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
