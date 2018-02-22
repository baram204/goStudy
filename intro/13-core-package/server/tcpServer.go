package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func 서버()  {
	// 포트 대기
	연결, 오류 := net.Listen("tcp",":9999")
	if 오류 != nil {
		fmt.Println(오류)
		return
	}
	for {
		// 연결수락
		접속, 오류 := 연결.Accept()
		if 오류 != nil {
			fmt.Println(오류)
			continue
		}

		// 연결처리
		go 서버접속처리(접속)
	}
}

func 서버접속처리(접속 net.Conn) {
	// 메시지 수신
	var 메시지 string

	// 접속을 가지고 메시지를 디코딩 해서 수신
	오류 := gob.NewDecoder(접속).Decode(&메시지)
	if 오류 !=nil {
		fmt.Println(오류)
	}else {
		fmt.Println("수신됨",메시지)
	}

	접속.Close()
}

func 클라이언트() {
	//서버에 연결
	접속, 오류 := net.Dial("tcp","127.0.0.1:9999")
	if 오류!=nil {
		fmt.Println(오류)
		return
	}
	// 메시지 송신
	메시지 := "안녕 고 서버야"
	fmt.Println("보냄", 메시지)

	// 접속 가지고 메시지를 인코딩해서 발송
	// gob 는 고 프로그램 끼리 값을 쉽게 주고받을 수 있게 한다.
	오류 = gob.NewEncoder(접속).Encode(메시지)
	if 오류 != nil {
		fmt.Println(오류)
	}

	접속.Close()
}

func main() {
		go 서버()
		go 클라이언트()

		// 시간 지연용 입력 함수
		var 입력 string
		fmt.Scanln(&입력)
	}
