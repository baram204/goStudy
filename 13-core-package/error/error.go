package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	오류 := errors.New("자체적인 오류 메시지")
	fmt.Println(오류)
}

