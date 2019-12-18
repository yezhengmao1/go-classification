package main

import (
	"classific"
	"data"
	"fmt"
	"util"
)


func GetClassificName(k int) string {
	var classifiname string = ""
	switch k {
	case 0:
		classifiname = "朴素贝叶斯"
	case 1:
		classifiname = "决策树    "
	}
	return classifiname
}

func main() {
	//dataPath := os.Args[1]
	//setnum, _ := strconv.Atoi(os.Args[2])

	setnum := 10
	/*
		wine  0
		bank  16
	*/
	dataPath := "/Users/yezm/Code/go-classification/data/"
	//classcol := 0
	//classcol := 16
	classcol := 4

	//alldata, colType := data.ReadWineDataToMatrix(dataPath + "wine.data")
	//alldata, colType := data.ReadBankDataToMatrix(dataPath + "bank.csv")
	alldata, colType := data.ReadIrIsDataToMatrix(dataPath + "iris.data")

	dataset := util.RandSplit(alldata, setnum)

	var pre [][]int = make([][]int, 3)
	var ana map[int][]float64 = make(map[int][]float64)
	ana[0] = make([]float64, 3) // 0 - 朴素贝叶斯
	ana[1] = make([]float64, 3) // 1 - 决策树

	for i := 0; i < setnum; i++ {
		testset := dataset[i]			// 测试集
		trainset := make([][]int, 0) 	// 训练集
		for j := 0; j < setnum; j++ {
			if i == j {
				continue
			}
			trainset = append(trainset, dataset[j] ...)
		}
		// 预测值
		pre[0] = classific.NavieBayes(trainset, testset, classcol)
		pre[1] = classific.DecisionTree(trainset, testset, classcol, colType)
		// 测试值
		var t []int = make([]int, 0)
		for _, item := range testset {
			t = append(t, item[classcol])
		}

		for k := 0; k < 2; k++ {
			f1,p,r := classific.F1(pre[k], t)

			fmt.Printf("%s: 第%d轮 训练集:%d 测试集:%d - 正确率%.2f 召回率%.2f F1:%.2f\n",
						GetClassificName(k) ,i, len(trainset), len(testset), p, r, f1)
			ana[k][0] += f1
			ana[k][1] += p
			ana[k][2] += r
		}
	}
	for k := 0; k < 2; k++ {
		fmt.Printf("%s 平均F1:%.4f 召回率:%.4f 正确率:%.4f\n",
			GetClassificName(k), ana[k][0]/float64(setnum), ana[k][2]/float64(setnum), ana[k][1]/float64(setnum))
	}
}
