package main

import (
	"os"
	"fmt"
	"log"
	"path/filepath"
	"io/ioutil"
)

func main() {

	print("번호를 선택해 주세요.")
	// 오류가 있다면 처리

	// 선택 변수 선언 후 Scan 함수로 입력받은 값을 저장하기 위해서
	// 포인터 주소를 넘겨준다.
	var 선택 int
	선택함수, 오류 := fmt.Scan(&선택)

	if 오류 != nil {
		return
	}

	// 선택을 입력받기
	println("실행 선택", 선택함수)

	switch 선택 {
	case 1:
		fileOpenFist()
	case 2:
		fileOpenSecond()
	case 3:
		createFile()
	case 4:
		getAllFilesNames()
	case 5:
		getAllFileNamesInDir()
	case 6:
		getFolderPathRecurcively()
	}

}

func fileOpenFist() {

	// 특정 경로의 파일 읽기
	// https://stackoverflow.com/questions/17071286/how-can-i-open-files-using-relative-paths-in-go
	파일, 오류 := os.Open(작업폴더() + "/src/intro/13-core-package/fileAndFolder/text.txt")
	// 고의 nil 은 자바의 null. 즉, 아무것도 참조하지 않는다는 값.
	// 스택에 생성된 것이 없고, 그 스택을 참조하지 않음을 나타냄
	// http://plaboratory.org/archives/5661
	오류처리(오류, "파일이 없습니다")

	// 같은 스콥에서 모든 것이 실행되고 난 뒤에 실행됨. 즉 무슨 짓이건 다 한 다음에 파일이 닫힌다.
	defer 파일.Close()

	// 파일 크기 구하기
	수치, 오류 := 파일.Stat()
	오류처리(오류, "파일과 연관된 수치를 알 수 없습니다.")

	// 파일을 읽음
	// 슬라이스를 만드는 것. 길이를 모르기 때문에 생성 시 길이(크기)를 알려준다.
	바이트 := make([]byte, 수치.Size())
	// 파일을 읽어서 넘겨준 바이트에 변수 공간에 담는다.
	_, 오류 = 파일.Read(바이트)
	오류처리(오류, "파일을 읽어서 변수에 담는 것을 실패했습니다.")

	str := string(바이트)
	fmt.Println(str)
}

func fileOpenSecond() {
	바이트, 오류 := ioutil.ReadFile(작업폴더() + "/src/intro/13-core-package/fileAndFolder/text.txt")
	오류처리(오류,"파일을 읽어서 바이트로 변환하는 데 실패했습니다.")
	문자열 := string(바이트)
	fmt.Println(문자열)
}

func createFile() {
	파일, 오류 := os.Create(작업폴더() + "/src/intro/13-core-package/fileAndFolder/createdText.txt")
	오류처리(오류, "파일 생성에 실패했습니다.")
	defer 파일.Close()

	파일.WriteString("생성된 파일 입니다.")

}

func 작업폴더() string {
	// 작업 디렉토디를 얻어온다.
	// https://gist.github.com/arxdsilva/4f73d6b89c9eac93d4ac887521121120
	dir, 오류 := os.Getwd()
	오류처리(오류, "")
	return dir
}


func 오류처리(오류 error, 메시지 string) {

	switch 오류 !=nil {
	case true:
			fmt.Println(메시지)
	default:
			return
	}
}

// 사용할 일이 없음
func getAllFilesNames() {

	// 파일을 얻어온다.
	// https://stackoverflow.com/questions/14668850/list-directory-in-go
	files, err := filepath.Glob("*")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files) // contains a list of all files in the current directory

	filess, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range filess {
		fmt.Println(f.Name())
	}
}

func getAllFileNamesInDir() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	파일정보들, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, 파일정보하나 := range 파일정보들 {
		fmt.Println(파일정보하나.Name())
	}
}

func getFolderPathRecurcively() {

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
