package main

import (
	"gwp/Chapter_2/data"
	"net/http"
)

func 인덱스(작성자 http.ResponseWriter, 요청 *http.Request) {
	쓰레드, 오류 := data.Threads()
	if 오류 != nil {
		error_message(작성자, 요청, "쓰레드 가져올 수 없음")
	} else {
		_, 오류 := 세션(작성자,요청)
		if 오류 != nil {
			HTML생성(작성자,쓰레드,"layout", "public.navbar", "index")
		} else {
			HTML생성(작성자,쓰레드,"layout", "public.navbar", "index")
		}
	}
}