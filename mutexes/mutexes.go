package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
	In the previous example we saw how to manage simple counter state using atomic operations.
	For more complex state we can use a mutex to safely access data across multiple goroutines.
	이전 예제에서 우리는 atomic 연산을 이용해서 간단한 카운터 상태를 어떻게 관리하는지 봤다.
	더 복잡한 상태를 위해서 우리는 다중 고루틴들 사이에 데이터에 안전하게 접근하기 위해서 mutex 를 사용할 수 있다.
*/

func main() {
	/*
		For our example the state will be a map.
		우리의 예제를 위해서 상태는 map이 될 것이다.
	*/
	var state = make(map[int]int)

	/*
		This mutex will synchronize access to state.
		이 mutex 는 상태에 동기적으로 접근할 것이다.
	*/
	var mutex = &sync.Mutex{}

	/*
		We'll keep track of how many read and write operations we do.
		우리는 얼마나 많은 읽기 및 쓰기 작업을 하는지 계속 추적할 것이다.
	*/
	var readOps uint64
	var writeOps uint64

	/*
		Here we start 100 goroutines to execute repeated reads against the state, once per millisecond in each goroutine.
		여기서 100개의 고루틴을 시작하여 각 고루틴에서 1ms당 한번씩 반복적인 읽기를 실행한다.
	*/
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				/*
					For each read we pick a key to access, Lock() the mutex to ensure exclusive access to the state, read the value at the chosen key, Unlock() the mutex, and increment the readOps count.
					각 읽기에서 우리는 접근을 위해 키를 하나 선택하고, Lock() 은 mutex가 상태로의 배타적인 접근과 선택된 키에서 값을 읽는 것을 보장하고 Unlock()은 mutex가 readOps 카운트를 증가시킨다.
				*/
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				/*
					Wait a bit between reads.
					읽는 것들 사이에 약간 기다린다.
				*/
				time.Sleep(time.Millisecond)
			}
		}()
	}

	/*
		We'll also start 10 goroutines to simulate writes, using the same pattern we did for reads.
		우리는 또한 쓰기를 가정한 10개의 고루틴들을 시작하고 읽기를 위해 사용했던 것과 같은 패턴을 사용한다.
	*/
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	/*
		Let the 10 goroutines work on the state and mutex for a second.
		state 와 mutex 상에서 작업할 고루틴들을 1초 동안 자게 한다.
	*/
	time.Sleep(time.Second)

	/*
		Take and report final operation counts.
		최종 작업 횟수를 파악하여 보고한다.
	*/
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps: ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps: ", writeOpsFinal)

	/*
		With a final lock of state, show how it ended up.
		상태의 마지막 잠금으로 어떻게 끝나는지 보여준다.
	*/
	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()

	/*
		Running the program shows that we executed about 90,000 total operations against our mutex-synchronized state.
		이 프로그램을 실행하면 우리가 mutex와 동기화된 상태에 대해 약 9만 건의 총 작업을 실행했다는 것을 알 수 있다.
	*/
}
