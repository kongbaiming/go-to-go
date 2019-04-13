package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	fmt.Println(len(os.Args)) //os.Args 启动时传递的参数 ，os.Args[0] 它本身
	fmt.Println(os.Args[0], os.Args[1])
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
