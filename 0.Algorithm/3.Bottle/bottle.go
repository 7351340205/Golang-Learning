package main

import "fmt"

func main() {
	a := bottle(20)

	fmt.Println(a)
}

/**
3个空瓶换1瓶饮料
递归解法
*/

func bottle(num int) int {
	if num < 3 {
		return 0
	}
	return num/3 + bottle(num/3)
}
