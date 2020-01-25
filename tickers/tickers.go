package main

import (
	"fmt"
	"time"
)

/*
	Timers are for when you want to do something once in the future - tickers are for when you want to do something repeatedly at regular intervals.
	Here's an example of a ticker that ticks periodically until we stop it.
	타이머는 여러분이 미래에 한 번 무언가를 하고 싶을 때를 위한 것이다. 티커는 여러분이 규칙적으로 반복적으로 무언가를 하고 싶을 때를 위한 것이다.
	여기 우리가 멈출 때까지 주기적으로 똑딱거리는 똑딱이의 예가 있다.
*/

func main() {
	/*
		Tickers use a similar mechanism to timers: a channel that is sent values.
		Here we'll use the select builtin on the channel to await the values as they arrive every 500ms.
		Ticker는 timer와 비슷한 매커니즘으로 사용한다: 값을 보내는 채널.
		우리는 채널에 내장된 select 를 사용해서 매 500ms 마다 값들이 도착하는 걸 기다린다.
	*/
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at ", t)
			}
		}
	}()
	/*
		Tickers can be stopped like timers.
		Once a ticker is stopped it won't receive any more values on its channel.
		We'll stop ours after 1600ms.
		Ticker는 timer가 멈추는 것과 같이 멈출 수 있다.
		한번 ticker가 멈추면 이 채널에서는 더이상 값을 받지 않는다.
		우리는 1600ms 뒤에 ticker를 멈출 것이다.
	*/
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
	/*
		When we run this program the ticker should tick 3 times before we stop it.
		우리가 이 프로그램을 실행시킬 때 ticker는 우리가 ticker를 멈추기 전까지 3번 tick할 것이다.
	*/
}
