package main

import ("fmt"; "container/list")

func main() {
	var x list.List
	x.PushBack("일")
	x.PushBack(2)
	x.PushBack("3")

	// 리스트 x 에서 첫 원소를 꺼낸 다음에
	// 원소가 참조할 것이 없을 때 까지 <- 이런식으로 not 표현이 생소하다.
	// 다음 원소로 넘어간다.
	for 원소 := x.Front(); 원소 != nil; 원소= 원소.Next() {
		fmt.Println(원소.Value)
		// 흐음 이게 타입 어셜션으로 변환을 하는 거구만요.
		// https://thebook.io/006806/ch04/04/04_01/
		결과, 성공 := 원소.Value.(int)
		if  성공 !=false {fmt.Println(결과, "변환 성공")} else {fmt.Println("정수형변환 실패")}

	}
}
