package main

import (
	"fmt"
	"time"
)

/*
	Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service.
	Go elegantly supports rate limiting with goroutines, channels, and tickers.
	Rate limiting 은 자원 활용 통제와 품질을 유지하기 위한 중요한 매커니즘이다.
	Go는 rate limiting을 고루틴, 채널 그리고 ticker 함께 우아하게 지원한다.
*/

func main() {
	/*
		First we'll look at basic rate limiting.
		Suppose we want to limit our handling of incoming requests.
		We'll serve these requests off a channel of the same name.
		첫번째로 우리는 기본적인 rate limiting 을 볼 것이다.
		우리는 들어오는 요청들의 핸들링을 제한하길 원한다고 가정하자.
		우리는 이 요청들을 같은 이름의 채널에서 처리할 것이다.
	*/
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	/*
		This limiter channel will receive a value every 200milliseconds.
		This is the regulator in our rate limiting scheme.
		이 limter 채널은 매 200ms 마다 값을 받을 것이다.
		이 것은 우리의 rate limiting 스키마 내에 regulator이다.
	*/
	limiter := time.Tick(200 * time.Millisecond)

	/*
		By blocking on a receive from the limiter channel before serving each request, we limit ourselves to 1 request every 200 milliseconds.
		각 요청을 전달하기 전에 limiter 채널에서 수신하는 것을 차단함으로써, 우리는 200ms마다 1개의 요청으로 제한한다.
	*/
	for req := range requests {
		<-limiter
		fmt.Println("request ", req, time.Now())
	}

	/*
		We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit.
		We can accomplish this by buffering our limiter channel.
		This burstyLimiter channel will allow bursts of up to 3 events.
		우리는 전반적인 rate limit을 유지하면서 우리의 rate limit 계획에서 짧은 요청 폭발을 허용하기를 원할 수 있다.
		우리는 이것을 우리의 limiter 채널을 buffer하는 것에 의해 달성할 수 있다.
		이 burstyLimiter 채널은 3개의 event들만 폭발하는 것을 허락할 것이다.
	*/
	burstyLimiter := make(chan time.Time, 3)

	/*
		Fill up the channel to represent allowed bursting.
		허용된 burst를 표현하기 위해서 채널을 채우자.
	*/
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	/*
		Every 200milliseconds we'll try to add a new value to burstyLimiter, up to its limit of 3.
		매 200ms마다 우리는 burstyLimiter 에 최대 3개의 새로운 값을 추가하기를 시도할 것이다.
	*/
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	/*
		Now simulate 5 more incoming requests.
		The first 3 of these will benefit from the burst capability of burstyLimiter.
		이제 요청이 5개 들어오는 것을 가정하자.
		이들 중 처음 3개는 bustyLimiter의 burst 수용으로 부터 이익을 얻을 것이다.
	*/
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request ", req, time.Now())
	}

	/*
		Running our program we see the first batch of requests handled once every ~200ms as desired.
		프로그램을 실행하면 200ms마다 한 번씩 처리되는 첫 번째 요청을 원하는 대로 볼 수 있다.
	*/

	/*
		For the second batch of requests we serve the first 3 immediately because of the burstable rate limiting, then serve the remaining 2 with ~200ms delays each.
		두 번째 요청의 경우 버스트 가능 속도 제한 때문에 첫 번째 3개 요청을 즉시 처리한 후 나머지 2개 요청을 각각 200ms씩 지연시킨다.
	*/
}
