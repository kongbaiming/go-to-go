// 查看某个变量的数据类型
// 查看变量占用的字节数
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 899999999
	fmt.Printf("变量a的类型是:%T,占用的字节数是：%d\n", a, unsafe.Sizeof(a))
	var price float32 = -8.999
	var price1 float64 = -8.99999456789
	fmt.Println("price:", price)
	fmt.Println("price1:", price1)
	var b byte = 'a'
	var c int = '我'
	fmt.Println(b)
	fmt.Printf("%c", b)
	fmt.Println()
	fmt.Printf("c的值是:%c,码值是:%d", c, c)
	fmt.Println()
	fmt.Printf("c:%T,占用的字节数是：%d\n", c, unsafe.Sizeof(c))
	fmt.Println()
	var d  int =  25106
	fmt.Printf("%c", d)
	// 字符可以运算，按照 码值运算
	fmt.Println()
	var n1 = 10 + 'a'
	fmt.Printf("c的值是:%c,码值是:%d", n1, n1)


}
