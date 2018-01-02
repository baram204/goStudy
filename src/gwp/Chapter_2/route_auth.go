package main

import (
	"data"
	"fmt"
	"net/http"
)

// 인증 관련 핸들러
// 그래서 핸들러펑션 인터페이스 함수 타입을 구현한다.

// Get /login
// 로그인 페이지 보이기
func 로그인(작성자 http.ResponseWriter, 요청 *http.Request) {
	// 템플릿 파일 이름들을 받아서 해석한 뒤 결과를 얻어온다.
	템플릿 := 템플릿파일들해석("login.layout", "public.navbar", "login")
	// 이것을 responseWriter 에다가 쓴다.
	템플릿.Execute(작성자, nil)
}

// GET /signup
// Show the signup page
func 가입(작성자 http.ResponseWriter, 요청 *http.Request) {
	HTML생성(작성자, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func 계정가입(작성자 http.ResponseWriter, 요청 *http.Request) {
	fmt.Println("계정가입 진입1")
	오류 := 요청.ParseForm()
	if 오류 != nil {
		danger(오류, "폼을 해석할 수 없습니다.")
	}
	fmt.Println("계정가입 진입2")
	용자 := data.User{
		Name:     요청.PostFormValue("이름"),
		Email:    요청.PostFormValue("이메일"),
		Password: 요청.PostFormValue("암호"),
	}

	fmt.Println(용자) // uuid 제외하고 모두 출력이 잘 된다.

	fmt.Println("계정가입 진입3")
	if err := 용자.Create(); err != nil {
		danger(err, "사용자를 생성할 수 없음")
	}
	fmt.Println("계정가입 진입4")
	http.Redirect(작성자, 요청, "/login", 302)
}
