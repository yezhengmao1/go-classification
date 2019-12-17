package util

import (
	"math/rand"
)

func randInt64(min, max int64) int64 {
	if min >= max || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

func shuffle(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		p := randInt64(0, int64(i))
		a := arr[i]
		arr[i] = arr[p]
		arr[p] = a
	}
	return arr
}

/* 随机将数据分为setnum组 */
func RandSplit(data [][]int, setnum int) [][][]int {
	var ret [][][]int = make([][][]int, setnum)
	var seq []int = make([]int, len(data))
	/* 每组个数 */
	var pre int = (len(data) + setnum - 1) / setnum
	var cnt int = 0
	for i := 0; i < len(data); i++ {
		seq[i] = i
	}
	seq = shuffle(seq)
	for i := 0; i < len(data); i++ {
		if i != 0 && i % pre == 0 {
			cnt++
		}
		ret[cnt] = append(ret[cnt], data[seq[i]])
	}
	return ret
}
