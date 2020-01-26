package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	The primary mechanism for managing state in Go is communication over channels.
	We saw this for example with worker pools.
	There are a few other options for managing state though.
	Here we'll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
	Go 내에서 상태를 관리하는 것을 위한 주요한 매커니즘은 채널들을 통해서 커뮤니케이션 하는 것이다.
	우리는 이 예제를 worker pool들에 대한 예제를 봤다.
	그러나 상태를 관리하기 위한 몇 가지 다른 선택사항이 있다.
	여기 다중 고루틴들에 의해 접근되어지는 atomic counter를 위한 sync/atomic 패키지 사용하는 예제를 볼 것이다.
*/

func main() {
	/*
		We'll use an unsigned integer to represent our(always-positive) counter.
		우리는 (항상 양수인) 카운터를 표현하기 위해서 부호없는 정수를 사용할 것이다.
	*/
	var ops uint64

	/*
		A WaitGroup will help us wait for all goroutines to finish their work.
		WaitGroup은 모든 고루틴들에 대해 그들의 작업이 끝나기를 기다리는 것을 도와줄 것이다.
	*/
	var wg sync.WaitGroup

	/*
		We'll start 50 goroutines that each increment the counter exactly 1000 times.
		우리는 카운터를 정확히 1000배 늘리는 50개의 고루틴들을 작성할 것임.
	*/
	for i := 0; i < 100; i++ {
		wg.Add(1)
		/*
			To atomically increment the counter we use AddUint64, giving it the memory address of our ops counter with the & syntax.
			카운터를 원자적으로 증가시키기 위해서 우리는 AddUint64를  사용해서 AddUint64에 &문법과 함께 우리의 ops 카운터의 메모리 주소를 줬다.
		*/
		go func() {
			for c := 0; c < 99999999; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	/*
		Wait until all the goroutines are done.
		모든 고루틴들이 끝나기를 기다린다.
	*/
	wg.Wait()

	/*
		It's safe to access ops now because we know no other goroutine is writing to it.
		Reading atomics safely while they are being updated is also possible, using functions like atomic.LoadUint64.
		이것은 이제 ops에 접근하기에 안전하다. 왜냐하면 우리는 다른 고루틴이 이것을 쓰지 않는다는 걸 알기 때문이다.
		atomic.LoadUint64 과 같은 기능을 이용해서 그들이 업데이트 되는 동안에도 atomic을 안전하게 읽을 수 있다.
	*/
	fmt.Println("ops: ", ops)
}
