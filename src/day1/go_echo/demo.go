// Struct 结构体  Array 数组 Slice 切片 Map 字典
package main

import "fmt"

var str string = "hello world"

const (
	a = iota + 50
	b
	c
	d, e, f = iota, iota, iota
	g       = iota
	h       = "h"
	i
	j = iota
)

const z = iota

func main() {

	fmt.Println(
		"a=", a,
		"b=", b,
		"c=", c,
		"d=", d,
		"e=", e,
		"f=", f,
		"g=", g,
		"h=", h,
		"i=", i,
		"j=", j,
		"z=", z)

}
