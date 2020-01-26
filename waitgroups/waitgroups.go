package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	To wait for multiple goroutines to finish, we can use a wait group.
	여러 고루틴이 끝나기를 기다리기 위해서 우리는 wait group을 사용할 수 있다.
*/

/*
	This is the function we'll run in every goroutine.
	Note that a WaitGroup must be passed to functions by pointer.
	이것은 모든 고루틴을 실행할 함수이다.
	WaitGroup은 반드시 포인터에 의해 함수로 전달되어야 한다.
*/
func worker(id int, wg *sync.WaitGroup) {
	/*
		On return, notify the WaitGroup that we're done.
		리턴에서 우리가 끝났다는 걸 WaitGroupo에 통지한다.
	*/
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	/*
		Sleep to simulate an expensive task.
		비싼 작업을 흉내내기 위해서 Sleep 한다.
	*/
	time.Sleep(time.Second * 2)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	/*
		This WaitGroup is used to wait for all the goroutines launched here to finish.
		이 WaitGroup은 이 곳에서 실행된 모든 고루틴들이 끝나기를 기다리는 데에 사용된다.
	*/
	var wg sync.WaitGroup

	/*
		Launch several goroutines and increment the WaitGroup counter for each.
		몇몇 고루틴들을 실행하고 WaitGroup 카운트를 각각 증가한다.
	*/
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	/*
		Block until the WaitGroup counter goes back to 0;
		all the workers notified they're done.
		WaitGroup 카운터가 0으로 돌아갈 때 까지 Block한다.
		모든 worker들은 그들이 끝날때 통지 받는다.
	*/
	wg.Wait()

	/*
		The order of workers starting up and finishing is likely to be different for each invocation.
		worker들의 시작하는 것과 끝나는 것의 순서는 각 호출마다 달라질 수 있다.
	*/
}
