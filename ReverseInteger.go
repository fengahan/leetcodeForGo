package main

import (
	"fmt"
)

func main() {
	v:=reverse(100)
	fmt.Println(v)
}
func reverse(x int) int {
	var  step  =0
	var  s=0
	for x!=0{
		e:=x%10
		x=x/10
		if step*10 > 2147483647{//正数溢出
			return 0
		}
		if step*10 < -2147483647 {//负数溢出
			return 0
		}
		step=step*10+e
	}
	if x<0 {
		return 0-s
	}
	return step
}
