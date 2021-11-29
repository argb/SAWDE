package main

import (
	"SAWDE/solution"
	"fmt"
)

func main() {
	arr := []int{7, 3, 7, 3, 1, 3, 4, 1}
	//arr := []int{2, 1, 1, 3, 2, 1, 1, 3}
	//arr := []int{7, 5, 2, 7, 2, 7, 4, 7}
	k := solution.GetK(arr)
	fmt.Println("K: ", k)
	span := solution.Solution(arr, k)
	fmt.Println("min span:", span)
	solution.ShowScript()
}
