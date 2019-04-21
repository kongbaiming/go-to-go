//字符串，固定长度的字符
// 字符串不可变
package main

import (
	"fmt"
)

func main() {
	var address string = "中关村大街"
	fmt.Printf("addres:%s\n",address)
	fmt.Println(address[2])
	str1 := "hello " +
		"world"
	str2 := "hello"
	str3 := "world"
	fmt.Println(str1)
	str4 := fmt.Sprintf("%s %s",str2,str3)
	fmt.Println(str4)
}
