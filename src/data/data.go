package data

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

/* 在此文件添加读取数据方法 */

func ReadWineDataToMatrix(path string) ([][]int, map[int]bool) {
	var ret [][]int = make([][]int, 0)
	var coltype map[int]bool = make(map[int]bool)

	strBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(strBytes), "\n")

	for i := 0; i < 14; i++ {
		coltype[i] = true
	}

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
	return ret, coltype
}

func ReadBankDataToMatrix(path string) ([][]int, map[int]bool) {
	var ret [][]int = make([][]int, 0)
	var coltype map[int]bool = make(map[int]bool)

	strBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(strBytes), "\n")

	var dict map[int]map[string]int = make(map[int]map[string]int)
	var indexArr map[int]bool = map[int]bool {
		1 : true,
		2 : true,
		3 : true,
		4 : true,
		6 : true,
		7 : true,
		8 : true,
		10 : true,
		15 : true,
		16 : true,
	}
	for i := 0; i < 17; i++ {
		coltype[i] = true
	}
	for k, _ := range indexArr {
		coltype[k] = false
	}
	/* 统计映射 */
	for _, item := range list[1:] {
		if len(item) < 3 {
			continue
		}
		attr := strings.Split(item, ";")
		for k, _ := range indexArr {
			if _, ok := dict[k]; !ok {
				dict[k] = make(map[string]int)
			}
			dict[k][attr[k]]++
		}
	}
	for _, v := range dict {
		pos := 0
		for n, _ := range v {
			v[n] = pos
			pos++
		}
	}

	for _, item := range list[1:] {
		if len(item) < 3 {
			continue
		}
		attr := strings.Split(item, ";")
		var itemVector []int = make([]int, 0)
		for i := 0; i < 17; i++ {
			if attr[1] == "unknown" || attr[2] == "unknown" {
				continue
			}
			if i == 0 || i == 5 || i == 9 || i == 11 || i == 12 || i == 13 || i == 14 {
				v, err := strconv.Atoi(attr[i])
				if err != nil {
					log.Fatal("no flag")
				}
				itemVector = append(itemVector, v)
			}else {
				itemVector = append(itemVector, dict[i][attr[i]])
			}
		}
		ret = append(ret, itemVector)
	}

	return ret, coltype
}

func ReadIrIsDataToMatrix(path string) ([][]int, map[int]bool) {
	var ret [][]int = make([][]int, 0)
	var coltype map[int]bool = make(map[int]bool)

	strBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(strBytes), "\n")

	for i := 0; i < 5; i++ {
		coltype[i] = true
	}

	for _, item := range list {
		if len(item) < 3 {
			continue
		}
		attr := strings.Split(string(item), ",")
		var itemVector []int = make([]int, 0)
		for i := 0 ; i < 4; i++ {
			v, _ := strconv.ParseFloat(attr[i], 64)
			itemVector = append(itemVector, int(v*10))
		}
		var c int
		if attr[4] == "Iris-setosa" {
			c = 0
		}else if attr[4] == "Iris-virginica" {
			c = 1
		}else if attr[4] == "Iris-versicolor" {
			c = 2
		}else {
			log.Fatal("参数错误")
		}
		itemVector = append(itemVector, c)
		ret = append(ret, itemVector)
	}
	return ret, coltype
}

func ReadMushRoomDataToMatrix(path string) ([][]int, map[int]bool) {
	var ret [][]int = make([][]int, 0)
	var coltype map[int]bool = make(map[int]bool)

	strBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(strBytes), "\n")

	var dict map[int]map[string]int = make(map[int]map[string]int)

	for i := 0; i < 23; i++ {
		coltype[i] = true
		dict[i] = make(map[string]int)
	}

	/* 统计映射 */
	for _, item := range list {
		if len(item) < 3 {
			continue
		}
		attr := strings.Split(item, ",")
		for k, _ := range attr{
			dict[k][attr[k]]++
		}
	}
	for _, v := range dict {
		pos := 0
		for n, _ := range v {
			v[n] = pos
			pos++
		}
	}

	/* 矩阵转化 */
	for _, item := range list {
		if len(item) < 3 {
			continue
		}
		attr := strings.Split(item, ",")
		var itemVector []int = make([]int, 0)
		for i := 0 ; i < 23; i++ {
			itemVector = append(itemVector, dict[i][attr[i]])
		}
		ret = append(ret, itemVector)
	}

	return ret, coltype
}

func ReadRoomDataToMatrix(path string) ([][]int, map[int]bool) {
	var ret [][]int = make([][]int, 0)
	var coltype map[int]bool = make(map[int]bool)
	strBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(strBytes), "\n")

	for i := 0; i < 6; i++ {
		coltype[i] = true
	}

	for _, item := range list[1:] {
		if len(item) < 3 {
			continue
		}

		attr := strings.Split(item, ",")
		var itemVector []int = make([]int, 0)

		for i := 2 ; i <= 7; i++ {
			var v int = 0
			t,_ := strconv.ParseFloat(attr[i], 64)
			if i == 2 || i == 3 {
				v = int(t + 0.5)
			}else if i == 4 || i == 5 || i == 7{
				v = int(t)
			}else {
				v = int(t * 10000)
			}
			itemVector = append(itemVector, v)
		}
		ret = append(ret, itemVector)
	}
	return ret, coltype
}