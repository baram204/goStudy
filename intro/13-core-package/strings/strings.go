package main

import (
	"fmt"
	ㄱ "strings"
)

func main() {

	var 문 = "김첨지"

	// 모든 타입을 받는 익명함수를 받아서 변수에 저장
	// https://newfivefour.com/golang-interface-type-assertions-switch.html
	ㅊ := func(t interface{}) {
		// 넘어온 인자의 타입에 따라서 분기후 실행
		switch v := t.(type) {

		case string, []string:
			fmt.Println(v)
		case bool:
			fmt.Println(v)
		case int:
			fmt.Println(v)
		case byte, []byte:
			fmt.Println(v)
		default:
			fmt.Println("unknown")
		}
	}

	ㅊ(ㄱ.Contains(문, "첨"))
	ㅊ(ㄱ.Count(문, "ㄱ"))
	ㅊ(ㄱ.HasPrefix(문, "김"))
	ㅊ(ㄱ.HasSuffix(문, "st"))
	ㅊ(ㄱ.Index(문, "첨"))
	ㅊ(ㄱ.Join([]string{"김", "첨"}, "판"))
	ㅊ(ㄱ.Repeat("판", 5))
	ㅊ(ㄱ.Split("기-이-임-처-엄-지-이", "-"))
	ㅊ(ㄱ.ToLower("KIMCHOMGEE"))
	ㅊ(ㄱ.ToUpper("kimcheomgee"))

	// 유니코드를 다루기 위해 rune 타입으로 담기
	룬문 := []rune{'김','첨','지'}
	룬투문 := string(룬문)

	// 유니코드 문자를 byte 로 바로 하면 overflow 가 난다.
	// 그래서 rune 을 string 한 다음에 그걸 byte 로 바꾼다.
	ㅊ(string([]byte(룬투문)))
	ㅊ([]byte(룬투문))

}
