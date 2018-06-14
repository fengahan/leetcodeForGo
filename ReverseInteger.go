package main

import (
	"strconv"
	"bytes"
	"fmt"
)

func main() {
	v:=reverse(100)
	fmt.Println(v)
}
func reverse(x int) int {

	var intString []rune
	if x >0 {
		intString=[]rune(strconv.Itoa(x))
	}else {
		intString=[]rune(strconv.Itoa(-x))
	}
	intStringLen:=len(intString)
	b:=new(bytes.Buffer)
	var v  string
	for i:=intStringLen;i>0;i-- {
		v=string(intString[i-1])
		b.WriteString(v)
	}
	s:=b.String()
	newInt,_:=strconv.Atoi(s)
	if x>0 && newInt<2147483648{
		return newInt
	}else if x<0 && 0-newInt>-2147483648 {
		return 0-newInt
	}
	return 0
}
