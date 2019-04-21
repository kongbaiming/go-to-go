//获取 int类型变量num的内存地址，将num的地址赋给指针ptr，通过ptr去修改num的值
package main

import "fmt"

func main() {
	var num int = 20
	fmt.Printf("num的内存地址是:%v\n",&num)
	var ptr *int = &num
	fmt.Printf("ptr指向的内存地址是:%v\n",ptr)
	*ptr = 22 //修改num的值
	fmt.Printf("num的值是:%v\n",num)
}

