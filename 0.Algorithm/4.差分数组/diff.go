package main

import "fmt"

func main() {

	a := []int{0, 0, 0, 0, 0}
	res := cal(a, 0, 4, 3)
	res = cal(res, 0, 2, 2)
	res = cal(res, 0, 1, 8)
	fmt.Println(res)

}

//构建差分数组
func cal(nums []int, s int, e int, num int) []int {
	l := len(nums) + 1
	a := make([]int, l)
	a[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		a[i] = nums[i] - nums[i-1]
	}

	a[s] += num
	if e+1 < len(nums) {
		a[e+1] -= num
	}

	res := make([]int, len(nums))
	res[0] = a[0]
	for j := 1; j < len(nums); j++ {
		res[j] = res[j-1] + a[j]
	}
	return res

}

//@ 2 直接用切片加 for循环暴力遍历
//func cal(nums []int, s int, e int, num int) []int {
//	a := nums[s : e+1]
//	for i, _ := range a {
//		a[i] += num
//	}
//	return nums
//}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
