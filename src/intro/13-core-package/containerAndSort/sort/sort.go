package main

import ("fmt"; "sort")

// 영문으로 된 타입을 한글로 사용하기 위해서 타입 포함(has)

type 사람 struct {
	 이름 string
	 나이 int
}

// 이름으로는 사람 슬라이스
type 이름으로 []사람

// 맨 아래를 참조
func (this 이름으로) Len() int {
	return len(this)
}

func (this 이름으로) Less(i,j int) bool {
	return this[i].이름 < this[j].이름
}

func (this 이름으로) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type 나이로 []사람
func (this 나이로) Len() int {
	return len(this)
}
func (this 나이로) Less(i, j int) bool {
	return this[i].나이 < this[j].나이
}
func (this 나이로) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}


func main() {
	아이들 := []사람{
		{"책이구지요",4},
		{"잭이구지",8},
		{"잭",6},
		{"잭이구",1},
		{"잭이지라용께",3},
	}
	sort.Sort(이름으로(아이들))
	fmt.Println(아이들)
	sort.Sort(나이로(아이들))
	fmt.Println(아이들)
}


/* // 메소드 명이 Len(), Less(), Swap() 이어야 sort 를 쓸 수 있다.
   // sort 패키지에는 다음처럼 interface 가 정의되어 있고.
   // Sort 매서드의 인자는 반드시 이 interface 만 받기 때문이다.
	type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
 */

