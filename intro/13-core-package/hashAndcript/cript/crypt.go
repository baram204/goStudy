package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"crypto/sha1"
)
// 비 암호화 해시에 비해서, 역 변환이 힘들다.


func 해시가져오기(파일이름 string)([]byte, error) {
	바이트, 오류 := ioutil.ReadFile(작업폴더()+"/src/intro/13-core-package/hashAndcript/non-cript/"+파일이름)

	if 오류 != nil {
		fmt.Println("파일 읽기 오유")
		return []byte{}, 오류
	}
	해시 := sha1.New()
	해시.Write(바이트)
	// sha1 은 160비트를 생성한다. 고로 바이트 슬라이스를 이용한다. 네이티브 타입 중에 160 비트 숫자 표현할 수 있는 것은 없다.
	return 해시.Sum([]byte{}), nil

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
	fmt.Println(h1, h2, h1[0] == h2[0])
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
