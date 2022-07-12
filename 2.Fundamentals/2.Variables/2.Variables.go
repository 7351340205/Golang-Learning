package main

import (
	. "fmt"
)

//声明全局变量并初始化
var X = 1
var Y = 2

func main() {
	Println(X, Y)
	//标准声明变量
	var name string
	var age int
	var isOk bool
	Println(name, age, isOk)

	//声明变量的同时赋值
	var name1 = "小王子"
	var age1 = 18
	var name2, age2 = "田螺哥", 24
	Println(name1, age1, name2, age2)

	//类型推导，由编译器完成
	var name3 = "小王子"
	var age3 = 18
	Println(name3, age3)

	//批量声明变量
	var (
		a string
		b int
		c float32
		d bool
	)
	Println(a, b, c, d)

	//简短变量声明
	//每次使用简短变量声明，必须有新变量被声明
	//否则编译不通过
	/**
		错误案例：
			a := "888"
		上方批量声明中已经声明过a变量，此处必须声明新变量
		如：
			a, n := "999", "999"
	**/
	short1 := "123"
	short2 := 456
	Println(short1, short2)

	//匿名变量不占用命名空间，不会分配内存。
	//使用 _ 表示
	// _ 对应的值 654 则不分配内存
	Z, _ := 321, 654
	Println(Z)

	//常量
	const Pi = 3.1415926535

	//批量声明常量
	const (
		j = 1
		k = 2
		l = 3
	)
	Println(Pi, j, k, l)

	const (
		i = 300
		o //默认o, p 的值与i一致
		p
	)
	Println(i, o, p)

	//iota只能用于常量
	//多用于枚举值
	//默认从0开始计数
	const (
		t = iota  //iota = 0
		y = "123" //iota = 1
		u         //iota = 2
		h = iota  //iota = 3
		g         //iota = 4
		_         //iota = 5	5由_跳过
		r         //iota = 6
	)
	Println(t, y, u, h, g, r)

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
	)

	Println(KB, MB, GB, TB, PB)
}
