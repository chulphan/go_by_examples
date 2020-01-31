package main

import (
	"fmt"
	"os"
)

/*
	Defer is used to ensure that a function call is performed later in program's execution, usually for purposes of cleanup.
	defer is often used where e.g. ensure and finally would be used in other languages.
	Dfer 는 함수 호출이 프로그램의 실행 내에서 나중에 실행되어진다고 보장하기 위해 사용한다. 대게 비우는 것의 목적들을 위한 것이다.
	defer는 다른 언어에서 확실히 그리고 최종적으로 사용되는 경우에 종종 사용된다.
*/

func main() {
	/*
		Suppose we wanted to create a file, write to it, and then close when we're done.
		Here's how we could do that with defer.
		우리가 파일을 만들고, 그것을 쓰고, 우리가 할 일을 다 끝내면 닫기를 원한다고 가정하자.
		여기 우리는 defer 를 사용해서 그 일을 할 수 있다.
	*/

	/*
		Immediately after getting a file object with createFile, we defer the closing of that file with closeFile.
		This will be executed at the end of the enclosing function(main), after writeFile has finished.
		createFile과 함께 직접적으로 파일 객체를 얻고 난 후에 우리는 closeFile과 함께 파일을 닫는 것을 연기할 수 있다.
		이것은 writeFile으 끝난 후 main 함수의 끝에서 실행될 것이다.
	*/

	f := createFile("./defer.txt")
	defer closeFile(f)
	writeFile(f)

	/*
		Running the program confirms that the file is closed after being written.
		프로그램을 실행시키면 파일은 쓰여지고 난 뒤에 닫히는 것을 확인할 수 있다.
	*/
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

/*
	It's important to check for errors when closing a file, even in a deferred function.
	연기된 함수내에 있다할지라도, 파일을 닫을때 에러를 체크하는 것은 중요하다.
*/
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
