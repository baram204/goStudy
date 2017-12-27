package main

import (
	b "bytes"
	"fmt"
)

func main() {

	// 오 버퍼에 담았다가 꺼내면 구지 rune 타입으로 변환하지 않아도 한글을 바이트화 하고 다시 꺼낼 수 있네.

	var buf b.Buffer
	buf.Write([]byte("테스트"))
	fmt.Println(buf)
	fmt.Println(string(buf.Bytes()))
}
