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
	case 2:
		classifiname = "逻辑回归  "
	}
	return classifiname
}

func main() {
	//dataPath := os.Args[1]
	//setnum, _ := strconv.Atoi(os.Args[2])

	setnum := 10

	dataPath := "data/"
	//classcol := 0
	//classcol := 16
	//classcol := 4
	//classcol := 0
	classcol := 5

	//alldata, colType := data.ReadWineDataToMatrix(dataPath + "wine.data")
	//alldata, colType := data.ReadBankDataToMatrix(dataPath + "bank.csv")
	//alldata, colType := data.ReadIrIsDataToMatrix(dataPath + "iris.data")
	//alldata, colType := data.ReadMushRoomDataToMatrix(dataPath + "mushroom.data")
	alldata, colType := data.ReadRoomDataToMatrix(dataPath + "occupancy.txt")

	dataset := util.RandSplit(alldata, setnum)

	var pre [][]int = make([][]int, 3)
	var ana map[int][]float64 = make(map[int][]float64)
	ana[0] = make([]float64, 3) // 0 - 朴素贝叶斯
	ana[1] = make([]float64, 3) // 1 - 决策树
	ana[2] = make([]float64, 3) // 2 - 逻辑回归

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
		pre[2] = classific.LogisticRegression(trainset, testset, classcol, 0.01, 500)
		// 测试值
		var t []int = make([]int, 0)
		for _, item := range testset {
			t = append(t, item[classcol])
		}

		for k := 0; k < 3; k++ {
			f1,p,r := classific.F1(pre[k], t)

			fmt.Printf("%s: 第%d轮 训练集:%d 测试集:%d - 正确率%.2f 召回率%.2f F1:%.2f\n",
						GetClassificName(k) ,i, len(trainset), len(testset), p, r, f1)
			ana[k][0] += f1
			ana[k][1] += p
			ana[k][2] += r
		}
	}
	for k := 0; k < 3; k++ {
		fmt.Printf("%s: 平均F1:%.4f 召回率:%.4f 正确率:%.4f\n",
			GetClassificName(k), ana[k][0]/float64(setnum), ana[k][2]/float64(setnum), ana[k][1]/float64(setnum))
	}
}
