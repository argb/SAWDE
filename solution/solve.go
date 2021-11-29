package solution

import (
	"fmt"
	"math"
)

const (
	EXPAND = "EXPAND"
	MOV	   = "MOV"
	SHRINK = "SHRINK"

)
var lastOp string
var lIndex int
var rIndex int

var minSpan int = math.MaxInt
var script []string

// Solution /**
/*
slide window + state transition
 */
func Solution(arr []int, k int) int {

	l := len(arr)
	for rIndex< l {
		for rIndex < l && !isOk(arr, k){
			expandR()
		}
		currSpan := rIndex-lIndex
		minSpan = min(currSpan, minSpan)
		fmt.Printf("1,lIndex:%d, rIndex:%d, minSpan:%d\n", lIndex, rIndex, minSpan)

		move() // 移动，移动之后可能符合条件，也可能不符合，分别处理
		fmt.Println("moved")

		for isOk(arr, k) {
			currSpan := rIndex-lIndex
			minSpan = min(currSpan, minSpan)
			fmt.Printf("2,lIndex:%d, rIndex:%d, minSpan:%d\n", lIndex, rIndex, minSpan)
			shrinkL()
		}
	}
	return minSpan
}

func isOk(arr []int, k int) bool {
	m := make(map[int]int)
	l := len(arr)
	var rBound int
	if rIndex<l {
		rBound = rIndex
	}else {
		rBound = l-1
	}
	for i:=lIndex; i <= rBound; i++ {
		m[arr[i]]++
	}
	if len(m) == k {
		minSpan = min(minSpan, rBound-lIndex) + 1
		return true
	}else {
		return false
	}
}

func expandR() {
	lastOp = EXPAND
	rIndex++
	script = append(script, lastOp)
}

func move() {
	lastOp = MOV
	lIndex++
	rIndex++
	script = append(script, lastOp)
}

func shrinkL() {
	lastOp = SHRINK
	lIndex++
	script = append(script, lastOp)
}

func GetK(arr []int) int {
	m := removeDuplicates(arr)
	k := len(m)
	return k
}

func removeDuplicates(arr []int) map[int]int {
	m := make(map[int]int)
	for _,v := range arr {
		m[v]++
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}else {
		return b
	}
}

func ShowScript() {
	fmt.Println("script: ", script)
}