/*练习1: 包中变量引用
练习2: 列出两数相加等于n的组合*/

package main

import (
	"fmt"
	"learn_golang/src/day2/lianxiti/cal"
)

func main() {
	fmt.Printf("%s\n", cal.Name1)
	fmt.Printf("%s\n", cal.Name2)
	cal.Add(10)

}
