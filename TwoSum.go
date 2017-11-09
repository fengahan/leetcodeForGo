package main

import "fmt"

func main() {
	nums:=[]int{3,2,4}
	target := 6
	result:=twoSum(nums,target)
	fmt.Println(result)
}
func twoSum(nums []int, target int) []int {
	c:=len(nums)
	result:=[]int{}
	for i:=0;i<c;i++{
		for j:=1;j<i+1;j++{
			if nums[i]+nums[j]==target {
				result=append(result,j)
				result=append(result,i)
			}
		}

	}
	return result
}