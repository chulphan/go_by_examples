package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/*
	In the previous example we used explicit locking with mutexes to synchronize access to shared state across multiple goroutines.
	Another option is to use the built-in synchronization features of goroutines and channels to achieve the same result.
	This channel-based approach aligns with Go’s ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine.
	이전 예제에서 우리는 mutex 를 이용해서 여러 고루틴들을 가로질러 공유된 상태에 동기적으로 접근해서 explicit locking 을 했다.
	또 다른 선택은 고루틴의 내장된 동기 특징들과 채널들을 이용해서 같은 결과를 달성하는 것이다.
	채널 기반 접근 방식은 정확히 1개의 고루틴이 각각의 데이터를 소유하고 통신하여 메모리를 공유한다는 Go의 아이디어와 일치한다.
*/

/*
	In this example our state will be owned by a single goroutine.
	This will guarantee that the data is never corrupted with concurrent access.
	In order to read or write that state, other goroutines will send messages to the owning goroutine and receive corresponding replies.
	These readOp and writeOp structs encapsulate those requests and a way for the owning goroutine to respond.
	이 예제에서 우리의 상태는 하나의 고루틴에 의해 소유될 것이다.
	이것은 데이터가 다른 동시접근에 대해 충돌될 일이 없다는 것을 보장할 것이다.
	readOp과 writeOp struct 는 그러한 요청과 소유하는 고루틴이 응답하는 방법을 캡슐화 한다.
*/

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	/*
		As before we'll count how many operations we perform.
		그 전에 우리는 우리가 실행할 연산들이 얼마나 많은지 세어볼 것이다.
	*/
	var readOps uint64
	var writeOps uint64

	/*
		The reads and writes channels will be used by other goroutines to issue read and write requests, respectively.
		읽기 및 쓰기 채널은 다른 고루틴들이 각각 읽기 및 쓰기 요청을 발행하는 데에 사용할 것이다.
	*/
	reads := make(chan readOp)
	writes := make(chan writeOp)

	/*
		Here is the goroutine that owns the state, which is a map as in the previous example but now private to the stateful goroutine.
		This goroutine repeatedly selects on the reads and writes channels, responding to requests as they arrive.
		A response is executed by first performing the requested operation and then sending a value on the response channel resp to indicate success (and the desired value in the case of reads).
		여기 이 상태를 소유하고 있는 고루틴이 있는데, 이것은 이전의 예와 같이 map이지만 현재 이 고루틴의 개인 map이다.
		이 고루틴은 읽기 및 쓰기 채널에서 반복적으로 선택하며, 도착 시 요청에 응답한다.
		요청된 작업을 먼저 수행한 다음 응답 채널 resp에 값을 전송하여 성공(및 판독의 경우 원하는 값)을 표시함으로써 응답을 실행한다.
	*/
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	/*
		This starts 100 goroutines to issue reads to the state-owning goroutine via the reads channel.
		Each read requires constructing a readOp, sending it over the reads channel, and the receiving the result over the provided resp channel.
		이것은 100개의 고루틴들이 reads 채널을 통해 상태를 소유하고 있는 고루틴에 reads를 발행하기 시작한다.
		각 read는 readOp을 구성하여 reads 채널을 통해 전송하고 제공된 resp 채널을 통해 결과를 수신해야 한다.
	*/
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 100; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	/*
		Let the goroutines work for a second.
		goroutine 에게 일을 하게 한다.
	*/
	time.Sleep(time.Second)

	/*
		Finally, capture and report the op counts.
		마침내 op 카운트들을 캡처하고 기록한다.
	*/
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps: ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps: ", writeOpsFinal)

	/*
		Running our program shows that the goroutine-based state management example completes about 80,000 total operations.
		프로그램을 실행하면 고루틴 기반 상태관리 예제는 약 80,000번의 총 연산을 완료하는 것을 보여준다. (나느 10만번..)
	*/

	/*
		For this particular case the goroutine-based approach was a bit more involved than the mutex-based one.
		It might be useful in certain cases though, for example where you have other channels involved or when managing multiple such mutexes would be error-prone.
		You should use whichever approach feels most natural, especially with respect to understanding the correctness of your program.
		이 특별한 경우, 고루틴 기반 접근법은 mutex 기반 접근법보다 약간 더 포함되었다.
		예를 들어 다른 채널이 관련되거나 그러한 mutex를 여러 개 관리할 경우 오류가 발생하기 쉬운 경우에 유용할 수 있다.
		특히 프로그램의 정확성을 이해하는 데 있어 가장 자연스럽게 느껴지는 접근법을 사용해야 한다.
	*/
}
