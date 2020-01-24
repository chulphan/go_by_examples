package main

import "fmt"

/*
	Basic sends and receives on channels are blocking.
	However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.
	채널 상에서 기본 발신과 수신은 blocking 이다.
	하지만 우리는 non-blockng 발신, 수신, 그리고 non-blocking 다중방법 select들을 구현하기 위해서 select를 default 절과 함께 사용할 수 있다.
*/
func main() {
	messages := make(chan string)
	signals := make(chan bool)
	// go func() {
	// 	messages <- "hahahaha"
	// }()
	/*
		Here's a non-blocking receive.
		If a value is available on messages then select will take the <-messages case with that value.
		If not it will immediately take the default case.
		여기 논블로킹 수신이다.
		만약 메세지에 값이 사용가능하다면 select 는 <-messages 케이스를 그 값과 취할 것이다.
		만약 아니라면 이건 직접적으로 default 케이스를 취할 것이다.
	*/
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	/*
		A non-blocking send works similarly.
		Here msg cannot be sent to the messages channel, because the channel has no buffer and there is no receiver.
		Therefore the default case is selected.
		논블로킹 발신도 비슷하게 동작한다.
		msg는 messages 채널로 보내질 수 없다. 왜냐하면 그 채널은 buffer도 아니고 수신자도 없기 때문이다.
		그러므로 default 케이스가 선택된다.
	*/
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	/*
		We can use multiple cases above the default clause to implement a multi-way non-blocking select.
		Here we attempt non-blocking receives on both messages and signals.
		우리는 위에 default절을 여러 경우에 사용할 수 있다. 다중방법 논블로킹 select를 구현하기 위해서.
		여기 우리는 non-blocking 수신하기를 시도한다 모든 mesages와 signals 상에서
	*/
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
