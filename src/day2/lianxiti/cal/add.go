package cal

import "fmt"

var Name1, Name2 string = "test1", "test2"

func Add(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("%d+%d=%d\n", i, n-i, n)
	}
}
