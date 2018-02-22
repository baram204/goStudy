package main

import (
	"net/http"
	"fmt"
)

// 핸들러 인터페이스를 구현하는 구조체
type 내멋대로핸들러 struct{}

func (h *내멋대로핸들러) ServeHTTP (rsw http.ResponseWriter, rq *http.Request) {
	fmt.Fprintf(rsw, "안녕 내멋대로 핸들러야")
}

// 나중에 핸들러함수 타입을 통해 변환 시킬 일반 함수
func 임의의핸들러함수 (rsw http.ResponseWriter, rq *http.Request) {
	fmt.Fprintf(rsw, "안녕 임의의 핸들러 함수야")
}

func 로그(h http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func (rsw http.ResponseWriter, rq *http.Request){
		fmt.Println("로깅이 됐습니다.", rq)
		h.ServeHTTP(rsw,rq)
	})
}


func 인증 (h http.Handler) http.HandlerFunc{

	return http.HandlerFunc(func (rsw http.ResponseWriter, rq *http.Request){
		fmt.Println("인증처리합니다.")
		h.ServeHTTP(rsw,rq)
	})
}

func main() {

	서버 := &http.Server{
		Addr:    "0.0.0.0:8000",
	}
	// 이미 핸들러 타입이니 등록만
	http.Handle("/", 인증(로그(&내멋대로핸들러{})))
	// Handlefunc(임의의핸들러함수)는 핸들러 타입으로 변환
	http.HandleFunc("/about", 인증(로그(http.HandlerFunc(임의의핸들러함수))))

	서버.ListenAndServe()
}
