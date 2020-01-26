package main

import (
	"fmt"
	"time"
)

/*
	In this example we'll look at how to implement a worker pool using goroutines and channels.
	이 예제에서는 고루틴과 채널을 이용해서 worker pool 을 어떻게 구현하는지 볼 것임.
*/

/*
	Here's the worker, of which we'll run several concurrent instances.
	These workers will receive work on the jobs channel and send the corresponding results on 'results'.
	We'll sleep a second per job to simulate an expensive task.
	여기에 worker가 있고 worker는 여러 인스턴스를 동시에 실행할 것이다.
	이 worker들은 jobs 채널 상에서 work를 받을 것이고 'results' 상에서 결과들에 대응하는 것을 보낸다.
	우리는 비싼 업무를 흉내내기 위해서 각 job마다 1초간 재울 것이다.
*/

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	/*
		In order to use our pool of workers we need to send them work and collect their results.
		We make 2 channels for this.
		우리의 worker pool을 이용하기 위해서 그들에게 일을 보내고 그들의 결과를 수집할 필요가 있다.
		우리는 이걸 위해서 2개의 채널들을 만든다.
	*/
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	/*
		This starts up 3 workers, initially blocked because there are no jobs yet.
		이건 3개의 worker들로 시작하고 처음에는 job들이 아직 없기 때문에 block된다.
	*/
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	/*
		Here we send 5 jobs and then close that channel to indicate that's all the work we have.
		여기 5개의 job을 보내고 그 채널을 우리가 가진 모든 work를 나타내기 위해 닫는다.
	*/
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	/*
		Finally we collect all the results of the work.
		This also ensures that the worker goroutines have finished.
		An alternative way to wait for multiple goroutines is to use a WaitGroup.
		마지막으로 우리는 그 일의 결과를 모두 수집한다.
		이것은 또 worker들의 고루틴이 끝났음을 보장한다.
		여러 고루틴을 기다리는 방법은 WaitGroup를 사용하는 것이다.
	*/
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
