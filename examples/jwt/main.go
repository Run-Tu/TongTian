package main

import (
	"fmt"
	"dqq/go/basic/pkg/common/xjwt"
)

func main() {
	token, err := xjwt.CreateToken(1, 2, true, 1000)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(token)
}
