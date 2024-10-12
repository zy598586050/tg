package tgcmd

import "fmt"

type Command struct{}

func (c Command) Run() {
	fmt.Println("测试三方包")
}
