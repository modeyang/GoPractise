package main

import (
	"fmt"
	"github.com/modeyang/GoPractise/interfaces"
	"github.com/modeyang/GoPractise/times"
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

	times.Timestamp()
}