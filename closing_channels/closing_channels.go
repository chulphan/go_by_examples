package main

import "fmt"

/*
	Closing a channel indicates that no more values will be sent on it.
	This can be useful to communicate completion to the channel's receivers.
	채널을 닫는 것은 값을 더 이상 보내지 않음을 나타낸다.
	이것은 채널의 수신자들과의 커뮤니티를 끝내기 위해 유용하다.
*/
func main() {
	/*
		In this example we'll use jobs channel to communicate work to be done from the main() goroutine to a worker goroutine.
		When we have no more jobs for the worker we'll close the jobs channel.
		이 예제에서 우리는 jobs 채널을 main() 고루틴으로 부터 worker 고루틴에게 일이 끝남을 소통하기 위해 사용한다.
		jobs가 worker를 위한 jobs가 더 이상 없을 때 jobs 채널을 닫을 것이다.
	*/
	jobs := make(chan int, 5)
	done := make(chan bool)

	/*
		Here's the worker goroutine. It repeatedly receives from jobs with j, more := <-jobs.
		In this special 2-value form of receive, the more value will be false if jobs has been closed and all values in the channel have already been received.
		We use this to notify on done when we've worked all our jobs.
		여기 worker 고루틴이다. 이건 jobs로 부터 j 와 more := <-jobs 로 부터 반복적으로 수신한다.
		이건 특별한 수신자의 2-value 폼이고 more 값은 jobs가 닫히면 false가 될 것이고 채널 내에 모든 값들은 이미 수신됐다.
		우리는 우리의 모든 작업을 완료했을 때 done 에 알리기 위해 윗 줄에 설명한 것을 사용한다.
	*/
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	/*
		This sends 3 jobs to the worker over the jobs channel, then closes it.
		이것은 3개의 작업을 worker 를 통해 jobs 채널로 보낸 다음에 jobs 채널을 닫는다.
	*/
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
	/*
		We await the worker using the synchronization approach we saw earlier
		우리는 앞에서 봤던 동기 접근을 사용해서 worker 를 기다리게한다.
	*/
}
