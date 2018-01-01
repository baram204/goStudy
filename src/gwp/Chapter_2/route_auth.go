package main

import (
	"net/http"
)

// 인증 관련 핸들러

// Get /login
// 로그인 페이지 보이기
func 로그인(작성자 http.ResponseWriter, 요청 *http.Request) {
	// 템플릿 파일 이름들을 받아서 해석한 뒤 결과를 얻어온다.
	템플릿 := 템플릿파일들해석("login.layout", "public.navbar", "login")
	// 이것을 responseWriter 에다가 쓴다.
	템플릿.Execute(작성자,nil)
}

