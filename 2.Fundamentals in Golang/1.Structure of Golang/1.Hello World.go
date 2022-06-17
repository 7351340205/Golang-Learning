//定义一个包的名字
package main

//导入外部包，调用外部包中的函数
//若 import . "fmt"
//则在下方代码中调用其函数无需加包名
//Println("xxx")

//若 import _ "fmt"
//则无需使用包也可以编译执行
import "fmt"

//func 为声明函数关键字
//每一个包仅能有一个main函数、一个init函数
//加载顺序为 导入的包init->导入的包main->本包init->本包main

//OutPut:
//Package main inited
//Hello World!

func main() {
	fmt.Println("Hello World!")
}

func init() {
	fmt.Println("Package main inited")
}

//关于注释：
//Golang支持C/C++风格的//注释
/**
 * 也支持 JavaLike的块注释
 **/
