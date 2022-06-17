package main

import "fmt"

func main() {

	var nums []int
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3

	fmt.Println(twoSum(nums, 3))

}

//LeetCode 01 两数之和
//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//你可以按任意顺序返回答案。
//
//Golang中Map底层实现为HashMap
//使用HashMap的去重功能
//循环遍历数组，得出数组下标

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
