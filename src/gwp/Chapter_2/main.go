package main

import (
	//	"html/template"
	"net/http"
)

//func 인덱스(작성자 http.ResponseWriter, 요청 *http.Request){
//	파일들 := []string{"template/layout.html",
//		"template/navbar.html",
//		"template/index.html",}
//		템플릿들 := template.Must(template.ParseFiles(파일들...))
//		쓰레드들, 오류 := data.쓰레드들(); if 오류 == nil {
//			템플릿들.ExecuteTemplate(작성자,"레이아웃",쓰레드들)
//	}
//}

func main() {

	멀플 := http.NewServeMux() // 서브하는 멀티플렉서 새로 만들기
	파일들 := http.FileServer(http.Dir(작업디렉토리() + "/public"))
	멀플.Handle("/static/", http.StripPrefix("/static/", 파일들))

	// 멀티 플레서에 요청에 따른 응답을 서브하는 핸들러를 등록한다.
	멀플.HandleFunc("/", 인덱스)
	//멀플.HandleFunc("/err", 오류)

	멀플.HandleFunc("/login", 로그인)
	멀플.HandleFunc("/signup", 가입)
	멀플.HandleFunc("/signup_account", 계정가입)
	멀플.HandleFunc("/authenticate", 인증)
	//
	//멀플.HandleFunc("/new", 새쓰레드)
	//멀플.HandleFunc("/create", 쓰레드생성)
	//멀플.HandleFunc("/post", 쓰레드넣기)
	//멀플.HandleFunc("/read", 쓰레드읽기)

	서버 := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: 멀플,
	}
	서버.ListenAndServe()
}
