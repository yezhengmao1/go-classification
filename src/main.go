package main

import (
	"classific"
	"data"
	"fmt"
	"util"
)


func main() {
	//dataPath := os.Args[1]
	//setnum, _ := strconv.Atoi(os.Args[2])

	setnum := 10
	classcol := 0
	dataPath := "/Users/yezm/Code/go-classification/data/"

	alldata := data.ReadWineDataToMatrix(dataPath + "wine.data")
	dataset := util.RandSplit(alldata, setnum)

	ar := 0.0
	ap := 0.0
	af := 0.0
	for i := 0; i < setnum; i++ {
		testset := dataset[i]			// 测试集
		trainset := make([][]int, 0) 	// 训练集
		for j := 0; j < setnum; j++ {
			if i == j {
				continue
			}
			trainset = append(trainset, dataset[j] ...)
		}
		pre := classific.NavieBayes(trainset, testset, classcol)
		var t []int = make([]int, 0)
		for _, item := range testset {
			t = append(t, item[classcol])
		}
		f1,p,r := classific.F1(pre, t)
		fmt.Printf("第%d轮: 正确率%.2f 召回率%.2f F1:%.2f\n", i, p, r, f1)
		ar += r
		ap += p
		af += f1
	}
	fmt.Printf("平均F1:%.4f 召回率:%.4f 正确率:%.4f\n", af/float64(setnum), ar/float64(setnum), ap/float64(setnum))
}
