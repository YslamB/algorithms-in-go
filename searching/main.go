package main

import (
	"fmt"
	"searching/linary"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	result := linary.Search(arr, target)
	fmt.Println(result)
}
