package main

import (
	"net/http"
	"io"
	"os"
	"fmt"
)

func 안녕(응답 http.ResponseWriter, 요청 *http.Request)  {
	응답.Header().Set(
		"Content-Type",
		// 그냥 보내면 문자가 깨지기 때문에 utf-8 문자 선언을 헤더에 집어넣는다.
		`text/html;charset=utf-8`,
	)
	io.WriteString(
		응답,
		`<doctype html>
		<html>
    <head>
        <title>안녕</title>
    </head>
    <body>
        이것은 받은 요청에 응답하는 겁니다.
    </body>
</html>`,
	)
}

func main() {
	 http.HandleFunc("/hello",안녕)

	 //폴더 := 작업폴더()
	 폴더 := 작업폴더()+"/src/intro/13-core-package/fileAndfolder"
	 파일서버 := http.FileServer(http.Dir(폴더))

	 // 스택에디트 참조
	// https://stackoverflow.com/questions/27945310/why-do-i-need-to-use-http-stripprefix-to-access-my-static-files/47997908#47997908
	 http.Handle("/static/",http.StripPrefix("/static",파일서버))

	 http.ListenAndServe(":9000",nil)
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
