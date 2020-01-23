package main

import "fmt"

/*
	Channels are the pipes that connect concurrent goroutines.
	You can send values into channels from one goroutine and receive those values into another goroutine.
	채널들은 동시실행하는 고루틴들을 연결한 파이프이다.
	하나의 고루틴으로 부터 값들을 채널로 보낼 수 있고 그 값들을 다른 고루틴에서 받을 수 있다.
*/

func main() {
	/*
		Create a new channel with make(chan val-type).
		Channels are typed by the values they convey.
		make(chan: 변수 타입)으로 새로운 채널을 생성한다.
		채널은 그들이 전달하는 값들에 의해 타입이 정해진다.
	*/
	messages := make(chan string)

	/*
		Send a value into a channel using the channel <- syntax.
		Here we send "ping" to the 'messages' channel we made above, from a new goroutine.
		채널 <- 문법을 사용해서 채널로 값을 보낸다.
		우리는 새로운 고루틴으로 부터 "ping"을 위에서 만들어 놓은  'messages' 채널에 보낸다.
	*/
	go func() { messages <- "ping" }()

	/*
		The <-channel syntax receives a value from the channel.
		Here we'll receive the "ping" message we sent above and print it out.
		<-channel 문법은 채널로부터 값을 받는다.
		우리는 우리가 위에서 받은 "ping" 메세지를 받고 출력할 것이다.
	*/
	msg := <-messages
	fmt.Println(msg)
}
