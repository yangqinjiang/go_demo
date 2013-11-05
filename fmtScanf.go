package main

import (
	"fmt"
)

func main() {
	var num int
	for i := 0; i < 10; i++ {
		fmt.Printf("Debug:i:%d\n", i)
		fmt.Scanf("%d", &num)
		fmt.Println(num)
	}
}
