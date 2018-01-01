package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

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

func 작업디렉토리() (dir string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return
}
