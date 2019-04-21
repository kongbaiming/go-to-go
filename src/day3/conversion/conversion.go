// 基本类型转换为string
// string 转基本数据类型
package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("=================分割线==========================")
	fmt.Println("基本数据类型转string")
	num1 := 100
	bool1 := true
	num2 := 3.45678
	char := 't'
	str := strconv.Itoa(num1)
	fmt.Printf("转换后类型为:%T,值为:%q\n", str, str)
	str1 := fmt.Sprintf("%d", num1)
	fmt.Printf("转换后类型为:%T,值为:%q", str1, str1)
	fmt.Println("")

	str2 := strconv.FormatInt(int64(num1), 10)
	fmt.Printf("转换后类型为:%T,值为:%q", str2, str2)
	fmt.Println("")
	str3 := fmt.Sprintf("%t", bool1)
	fmt.Printf("转换后类型为:%T,值为:%q", str3, str3)
	str4 := strconv.FormatBool(bool1)
	fmt.Println("")
	fmt.Printf("转换后类型为:%T,值为:%q", str4, str4)
	str5 := fmt.Sprintf("%f", num2)
	fmt.Println("")
	fmt.Printf("转换后类型为:%T,值为:%q", str5, str5)
	str6 := strconv.FormatFloat(num2, 'f', 2, 64)
	fmt.Println("")
	fmt.Printf("转换后类型为:%T,值为:%q", str6, str6)
	str7 := fmt.Sprintf("%c", char)
	fmt.Println("")
	fmt.Printf("转换后类型为:%T,值为:%q", str7, str7)

	fmt.Println("")
	fmt.Println("=================分割线==========================")
	fmt.Println("string转基本数据类型")

	var new_str1 = "true"
	a, _ := strconv.ParseBool(new_str1)
	fmt.Printf("转换后类型为:%T,值为:%v\n", a, a)
	var new_str2 = "123456789"
	b ,_:= strconv.ParseInt(new_str2,10,64)
	fmt.Printf("转换后类型为:%T,值为:%v\n", b, b)
	var new_str3 = "123.45678"
	c,_ :=strconv.ParseFloat(new_str3,64)
	fmt.Printf("转换后类型为:%T,值为:%v\n", c, c)
}
