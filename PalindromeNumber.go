package main

import "fmt"

func main() {

	var n=10
	v:=Palindrome(n)
	fmt.Println(v)
}
func Palindrome(number int)bool  {
	if number<0 {
		return false
	}
	var n int
	n=number
	nv:=number
	var arr =make([]int,0)
	for nv >0{
		n=nv%10
		nv=nv/10
		arr=append(arr,n)
	}
	l:=len(arr)
	var f =true
	for i:=0;i<l ;i++  {
		if (arr[i]==arr[l-i-1]){
			continue
		}else {
			f=false
		}
	}

	return f
}
