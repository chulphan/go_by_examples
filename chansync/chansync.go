package main

import (
	"fmt"
)

/*
	We can use channels to synchronize execution across goroutines.
	Here’s an example of using a blocking receive to wait for a goroutine to finish.
	When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.
	우리는 채널들을 사용하여 고루틴들을 가로질러 동기적으로 실행시킬 수 있다.
	여기 예제는 고루틴이 끝날때까지 수신을 막을 수 있다.
	다수의 고루틴들이 끝날때까지 기다리고 있을때, 넌 아마 WaitGroup 을 사용하는 것을 선호할 수 있다.
*/

/*
	This is the function we'll run in a goroutine.
	The 'done' channel will be used to notify another goroutine that this function's work is done.
	이것은 고루틴 내에서 실행될 함수다.
	'done' 채널은 이 함수의 동작이 끝남을 다른 고루틴에게 알릴때 사용될 것이다.
*/

func worker(done chan bool, order int) {
	fmt.Printf("working...%d", order)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		// time.Sleep(time.Second)
	}
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 2)
	/*
		Start a worker goroutine, giving it the channel to notify on.
		worker 고루틴을 시작한다, worker 고루틴에게 알림을 받을 채널을 넘겨준다.
	*/
	go worker(done, 1)
	go worker(done, 2)
	/*
		Block until we receive a notification from the worker on the channel.
		채널 상에 worker로 부터 알림을 받을때까지 Block!
	*/
	<-done
	<-done

	/*
		If you removed the <- done line from this program, the program would exit before the worker even started.
		만약에 <-done 라인을 이 프로그램에서 삭제하면, 프로그램은 worker 가 실행 되기도 전에 끝날것임.
	*/
}
