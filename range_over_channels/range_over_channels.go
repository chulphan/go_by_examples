package main

import "fmt"

/*
	In a previous example we saw how for and range provide iteration over basic data structures.
	We can also use this syntax to iterate over values received from a channel.
	이전 예제에서 우리는 기본 데이터 구조에 대한 반복을 제공하는 방법과 범위를 보았다.
	우리는 또한 이 구문을 사용하여 채널에서 수신한 값을 반복할 수 있다.
*/
func main() {
	/*
		We'll iterate over 2 values in the queue channel.
		우리 queue 채널 내에 두 개의 값들을 순회할 것이다.
	*/
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	/*
		This range iterates over each element as it's received from queue.
		Because we closed the channel above, the iteration terminates after receiving the 2 elements.
		이 range는 queue로 부터 수신한 각 원소를 순회한다.
		왜냐하면 우리가 위에서 채널을 닫았기 때문에 두 원소를 받은 후에 순회를 종료한다.
	*/
	for elem := range queue {
		fmt.Println(elem)
	}

	/*
		This example also showed that it's possible to close a non-empty channel but still have the remaining values be received.
		이 예제는 또한 비어있지 않은 채널을 닫을 수 있지만 나머지 값을 여전히 수신할 수 있다는 것을 보였다.
	*/
}
