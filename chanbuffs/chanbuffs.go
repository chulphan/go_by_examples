package main

import "fmt"

/*
	By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
	Buffered channels accept a limited number of values without a corresponding receiver for those values.
	기본적으로 채널들은 unbuffered 인데 의미는 채널들은 오직 보낸 값을 받을 준비가 되어있는 수신 측이 있을 때만 발신을 받아들인다.
	버퍼 채널은 해당 값에 해당하는 수신기가 없는 제한된 수의 값을 허용한다.
*/

func main() {
	/*
		Here we make a channel of strings buffering up to 2 values.
		최대 2개의 값을 버퍼링 하는 문자열 채널을 만든다.
	*/
	messages := make(chan string, 2)

	/*
		Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
		이 채널이 buffered 이기 때문에, 우리는 이 값들을 대응하는 동시 수신 없이도 채널로 보낼 수 있다.
	*/
	messages <- "buffered"
	messages <- "channel"

	/*
		나중에 우리는 이 두 값들을 보통적인 방법으로 받을 수 있다.
	*/
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
