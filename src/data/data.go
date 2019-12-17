package data

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

/* 在此文件添加读取数据方法 */

func ReadWineDataToMatrix(path string) [][]int {
	var ret [][]int = make([][]int, 0)
	strBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(strBytes), "\n")
	for _, item := range list {
		if len(item) < 3 {
			continue
		}
		attr := strings.Split(string(item), ",")
		var itemVector []int = make([]int, 0)
		for i := 0 ; i < 14; i++ {
			v, _ := strconv.ParseFloat(attr[i], 64)
			if i == 0 || i == 1 || i == 4{
				itemVector = append(itemVector, int(v))
			}else if i == 5 || i == 13{
				itemVector = append(itemVector, int(v/10))
			}else {
				itemVector = append(itemVector, int(v * 10))
			}
		}
		ret = append(ret, itemVector)
	}
	return ret
}

func ReadBankDataToMatrix(path string) [][]int {
	var ret [][]int = make([][]int, 0)
	return ret
}

