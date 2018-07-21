package main

import "fmt"

func main() {

	var n=101101
	v:=Palindrome(n)
	fmt.Println(v)
}
func Palindrome(number int)bool  {
	if number<0 {
		return false
	}
	var n int

	var arr =make([]int,0)
	for number >0{
		n=number%10
		number=number/10
		arr=append(arr,n)
	}
	l:=len(arr)
	var f =true
	for i:=0;i<l ;i++  {
		if (arr[i]==arr[l-i-1]){
			continue
		}else {
			f=false
			break
		}
	}

	return f
}
