/*练习1: 包中变量引用
练习2: 列出两数相加等于n的组合*/

package main

import (
	"fmt"
	a "learn_golang/src/day2/lianxiti/cal"
)

func main() {
	a.Add(10)
	fmt.Printf("%s\n", a.Name1)
	fmt.Printf("%s\n", a.Name2)

}
