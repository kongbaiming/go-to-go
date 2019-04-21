//指针变量存的是一个地址，地址指向的空间才是值
package main

import "fmt"

func main() {
	var i int = 10
	fmt.Println("i的内存地址:",&i)
	// ptr 指针类型
	var ptr *int = &i
	fmt.Printf("ptr的值:%v\n",ptr)
	fmt.Printf("ptr的内存地址是:%v\n",&ptr)
	//取出指针指向的值
	j := *ptr
	fmt.Printf("ptr指向的值是:%v\n",j)

}
