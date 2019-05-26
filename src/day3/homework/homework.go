// 1.  假如还有97天放假，问，还剩几个星期零几天
// 2.  求华氏温度对应的摄氏温度

/*
摄氏温标（C）和华氏温标（F）之间的换算关系为：

F=C×1.8+32

C=(F-32)÷1.8
*/
package main

import (
	"fmt"
)

func test() bool {
	fmt.Println("test func")
	return true

}

func main() {
	var days int = 97
	weeks := days / 7
	day := days % 7
	fmt.Printf("还剩%d个星期%d天\n", weeks, day)

	fahrenheit := 143.2
	var celsiust float32 = float32(fahrenheit-32) / 1.8
	fmt.Printf("华氏度%v度换算成摄氏度后为:%v度\n", fahrenheit, celsiust)

	var i int = 10
	if i < 9 || test() {
		fmt.Println("HELLO")
	}

}
