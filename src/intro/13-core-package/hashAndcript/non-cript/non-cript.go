package main

import (
	"hash/crc32"
	"fmt"
	"io/ioutil"
	"os"
)

func 해시가져오기(파일이름 string)(uint32, error) {
	바이트, 오류 := ioutil.ReadFile(작업폴더()+"/src/intro/13-core-package/hashAndcript/non-cript/"+파일이름)

	if 오류 != nil {
		fmt.Println("파일 읽기 오유")
		return 0, 오류
	}
	해시 := crc32.NewIEEE()
	해시.Write(바이트)
	return 해시.Sum32(), nil

}

func main() {
	h1, err := 해시가져오기("text1.txt")
	if err != nil {
		return
	}
	h2, err := 해시가져오기("text2.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
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
