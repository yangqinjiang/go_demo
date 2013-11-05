package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9, 10, 11, 12, 13, 14}
	copy(s2[2:5], s1[1:4]) //s2为目的地,s1为源  s1->s2
	fmt.Println(s2)
}
