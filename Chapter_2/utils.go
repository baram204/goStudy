package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"gwp/Chapter_2/data"
	"errors"
)

var logger *log.Logger

func 세션(작성자 http.ResponseWriter, 요청 *http.Request) (ㅅㅅ data.Session, 오류 error) {
	 쿠키, 오류 := 요청.Cookie("_cookie")
	 if 오류 == nil {
	 	ㅅㅅ = data.Session{Uuid: 쿠키.Value}
	 	if 옼, _ := ㅅㅅ.Check(); !옼 {
	 		오류 = errors.New("유효하지 않은 세션")
		}
	 }
	return

}

// HTML 템플릿 파일 해석
// 파일 이름의 목록을 통과시켜 템플릿을 얻는다.
func 템플릿파일들해석(파일이름 ...string) (템 *template.Template) {

	var 파일들 []string
	템 = template.New("레이아웃")
	for _, 파일 := range 파일이름 {
		// 파일들 슬라이스 내부에다가, 파일이름과 템플릿경로 혼합 출력된 문자열을 추가.
		파일들 = append(파일들, fmt.Sprintf(작업디렉토리()+"/templates/%s.html", 파일))
	}
	// 해석된 결과를 처리하기 위해 Must 사용, Must 는 오류 검출도 해준다.
	템 = template.Must(템.ParseFiles(파일들...))
	return
}

// 넘겨준 자료와 파일 이름 그리고 자료를 가지고 템플릿을 생성한다.
func HTML생성(작성자 http.ResponseWriter, 자료 interface{}, 파일이름들 ...string) {
	var 파일들 []string
	for _, 파일 := range 파일이름들 {
		파일들 = append(파일들, fmt.Sprintf(작업디렉토리()+"/templates/%s.html", 파일))
	}

	템플릿들 := template.Must(template.ParseFiles(파일들...))
	템플릿들.ExecuteTemplate(작성자, "레이아웃", 자료)
}

func 작업디렉토리() (dir string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

// Convenience function to redirect to the error message page
func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}
