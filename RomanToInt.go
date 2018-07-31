package main

import (
	"fmt"
)

func main() {
	romV:="MCMXCIV"
	v:=romanToInt(romV)
	fmt.Println(v)
}
func romanToInt(s string) int {
	var romMap=make(map[string]int)
	var sln=len(s)
	romMap["I"]=1
	romMap["V"]=5
	romMap["X"]=10
	romMap["L"]=50
	romMap["C"]=100
	romMap["D"]=500
	romMap["M"]=1000
	var g,v,sum int
	if len(s)==1{
		return romMap[string(s[0])]
	}
	for i:=0;i<sln;i++ {
		if i+1>sln-1 {
			break
		}
		current:=romMap[string(s[i])]
		next:=romMap[string(s[i+1])]
		v=current
		if current<next {//
			g=next/current
			if  g==5 || g==10 {
				v=next-current
			}else {
				v=next
			}
			i=i+1
		}else if next==current{
			v=next+current
			if i+1>sln-2 {
				break
			}
			fullNext:=romMap[string(s[i+2])]
			if next==fullNext {
				v=v+ romMap[string(s[i+2])]
			}
		}
		sum=sum+v
	}
	return sum
}