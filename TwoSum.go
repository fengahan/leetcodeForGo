package main

import "fmt"

func main() {
	nums:=[]int{2,3,3,6,12}
	target := 8
	result:=twoSum(nums,target)
	fmt.Println(result)
}
func twoSum(nums []int, target int) []int {
	c:=len(nums)
	result:=[]int{}
	m:=make(map[int]int)
	for i:=0;i<c;i++{

		if _,ok:=m[nums[i]];ok && m[nums[i]]!=i{
			result=append(result,m[nums[i]])
			result=append(result,i)
		}else {
			m[target-nums[i]]=i
		}
	}
	return result
}