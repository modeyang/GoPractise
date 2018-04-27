package main

import (
	"fmt"
	"github.com/modeyang/GoPractise/interfaces"
	"github.com/modeyang/GoPractise/debug"
	"time"
	"github.com/modeyang/GoPractise/web"
)

//  for range 遍历数组时，是否遍历的是副本？请写程序验证一下
func ArrayIter() {
	a := [3]int{1, 2, 3}
	for _, item := range(a) {
		a[2] = 4
		fmt.Print(item)
	}
	fmt.Println("")
}


func main() {
	ArrayIter()
	var a interfaces.A = &interfaces.B{}
	fmt.Println(a.Haha(1))

	debug.SetupStackTrap()

	t := time.NewTimer(3 * time.Second)

	go func() {
		select {
		case <- t.C:
			debug.StartStack()
		}
	}()

	time.Sleep(10 * time.Second)
	//web.NewEndlessServer(5070)
}