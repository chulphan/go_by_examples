package main

import "os"

/*
	A panic typically means something went unexpectedly wrong.
	Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.
	panic은 통상적으로 무엇인가가  기대하지 않게 잘못되고 있는 것을 뜻한다.
	대개 우리는 정상 작동 중에 발생해서는 안 되는 오류나 우아하게 처리할 준비가 되어 있지 않은 오류에 빠르게 실패하기 위해 그것을 사용한다.
*/

func main() {
	/*
		We’ll use panic throughout this site to check for unexpected errors.
		This is the only program on the site designed to panic.
		우리는 이 사이트가 기대하지 않은 에러를 내는지 확인하기 위해서 이 사이트 전체에 panic을 사용할 것이다.
		이것은 그 사이트에서 공황상태에 빠지도록 고안된 프로그램이다.
	*/
	panic("a problem")

	/*
		A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle.
		Here’s an example of panicking if we get an unexpected error when creating a new file.
		패닉의 일반적인 사용은 함수에서 처리 방법을 알지 못하는 오류 값을 반환할 경우 중단하는 것이다.
		여기 예제는 새로운 파일을 만들때 예상치 못한 오류가 발생하면 panic 하는 예제이다.
	*/
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	/*
		Running this program will cause it to panic, print an error message and goroutine traces, and exit with a non-zero status.
		이 프로그램을 실행하면 panic 을 일으킬 것이고 에러 메세지와 고루틴 trace 그리고 exit with non-zero status 를 출력할 것이다.
	*/

	/*
		Note that unlike some languages which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values wherever possible.
		참고로 많은 에러를 처리하기 위해 예외를 사용하는 다른 언어들과는 다르게 Go에서는 가능한 경우 error를 표현하는 반환값을 사용하는 것이 관용이다.
	*/
}
