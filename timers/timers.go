package main

import (
	"fmt"
	"time"
)

/*
	We often want to execute Go code at some point in the future, or repeatedly at some interval.
	Go's built-in 'timer' and 'ticker' features make both of these tasks easy.
	We'll look first at timers and then at tickers.
	우리는 종종 Go 코드를 미래의 어느 시점 또는 어떠한 구간을 반복적으로 실행시키길 원한다.
	Go의 내장된 timer 와 ticker 특징은 이러한 작업들을 쉽게 할 수 있도록 만든다.
	우리는 처음으로 timers 를 본 뒤에 tickers 를 볼 거임.
*/

func main() {
	/*
		Timers represent a single event in the future.
		You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.
		This timer will wait 2seconds.
		Timer는 미래에 단독 이벤트를 표현한다.
		timer에게 얼마나 기다릴지를 말하면 그 시점에 알림을 받을 채널을 제공한다.
		이 timer는 2초를 기다릴 것이다.
	*/
	timer1 := time.NewTimer(2 * time.Second)

	/*
		The <-timer1.C blocks on the timer's channel C until it sends a value indicating that the timer field.
		<-timer1.C 는 timer의 채널 C를 timer 필드가 나타내는 값을 보내기 전까지 막는다.
	*/
	<-timer1.C
	fmt.Println("Timer 1 fired")

	/*
		If you just wanted to wait, you could have used time.Sleep.
		One reason a timer may be useful is that you can cancel the timer before it fires.
		Here's an example of that.
		만약 그냥 기다리기만 하고 싶으면 time.Sleep 를 사용할 수 있다.
		timer가 유용할 수 있는 한가지 이유는 timer 가 발사되기 전에 취소할 수 있다.
		여기 그 예제가 있다.
	*/
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 stopped.")
	}
	/*
		Give the timer2 enough time to fire, if it ever was going to, to show it is in fact stopped.
		타이머2를 발사할 수 있는 충분한 시간을 주어 실제로 정지되었음을 표시한다.
	*/
	time.Sleep(2 * time.Second)

	/*
		The first timer will fire ~2s after we start the program, but the second should be stopped before it has a chance to fire.
		첫번째 timer는 우리가 프로그램을 시작하고 2초 후에 발사되지만 두번째 timer는 발사되기 전에 멈춰진다.
	*/
}
