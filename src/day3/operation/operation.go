package main

import "fmt"

func main() {
	//如果运算数都是整数，结果保留整数部分 10 /4 =  2
	var a int
	var b int
	a = 10
	b = 3
	fmt.Println(a / b)
	var n1 float32 = 10 / 4 // n1 = 2
	fmt.Println(n1)
	// 如果希望保留小数部分，需要浮点数参与运算
	var n2 float32 = 10.0 / 4 // n2 = 2.5
	fmt.Println(n2)

	//  a % b = a - a/b *b
	fmt.Println(10 % 3)   // 1
	fmt.Println(-10 % 3)  //-1
	fmt.Println(10 % -3)  //1
	fmt.Println(-10 % -3) //-1

	
}
