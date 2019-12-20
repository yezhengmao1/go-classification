package classific

import (
	"log"
	"math"
	"math/rand"
)

/* 预测类型 */
type PredictNode struct {
	weightVector []float64
	/* 划分为两类预测 */
	classlabelL []int	// 0
	classlabelR []int	// 1
	lchild *PredictNode
	rchild *PredictNode
}

func (node *PredictNode) GetLable(v int) float64 {
	for _, i := range node.classlabelL {
		if i == v {
			return 0
		}
	}
	for _, i := range node.classlabelR {
		if i == v {
			return 1
		}
	}
	log.Fatal("错误")
	return -1
}

func (node *PredictNode) LogisticPredict(testcase []int, classcol int) int {
	var v float64 = 0.0
	var p int = 0
	for col, _ := range testcase{
		if col == classcol {
			continue
		}
		v += node.weightVector[p] * float64(testcase[col])
		p++
	}
	v = 1 / (1 + math.Exp(-v))
	if v >= 0.5 {
		if node.rchild == nil {
			if len(node.classlabelR) != 1 {
				log.Fatal("错误测试")
			}
			return node.classlabelR[0]
		}
		return node.rchild.LogisticPredict(testcase, classcol)
	}else {
		if node.lchild == nil {
			if len(node.classlabelL) != 1 {
				log.Fatal("错误测试")
			}
			return node.classlabelL[0]
		}
		return node.lchild.LogisticPredict(testcase, classcol)
	}
}

func NewPredictNode(classmap map[int]int, attrnum int) *PredictNode {
	if len(classmap) < 2 {
		log.Fatal("错误")
		return nil
	}

	cnt := 0
	split := len(classmap) / 2
	var labelL []int = make([]int, 0)
	var labelR []int = make([]int, 0)

	var weight []float64 = make([]float64, attrnum)
	for i := 0; i < attrnum; i++ {
		weight[i] = rand.Float64()
	}

	for k, _ := range classmap {
		if cnt >= split {
			labelR = append(labelR, k)
		}else {
			labelL = append(labelL, k)
		}
		cnt++
	}

	return &PredictNode{
		weightVector: weight,
		classlabelL:  labelL,
		classlabelR:  labelR,
		lchild:       nil,
		rchild:       nil,
	}
}

func DeleteLabelInSet(dataset [][]int, classcol int, deletelabel []int) [][]int{
	var ret [][]int = make([][]int, 0)
	var todelete map[int]bool = make(map[int]bool)

	for _, v := range deletelabel {
		todelete[v] = true
	}
	for _, item := range dataset {
		if _, ok := todelete[item[classcol]]; !ok {
			ret = append(ret, item)
		}
	}
	return ret
}

func LogisticRegressionTrain(dataset [][]int, classcol int, alpha float64, iter int) *PredictNode {
	var classmap map[int]int = make(map[int]int)
	for _, item := range dataset {
		classmap[item[classcol]]++
	}

	root := NewPredictNode(classmap, len(dataset[0]) - 1)

	for i := 0; i < iter; i++ {
		for _, item := range dataset {
			/* 计算权重 */
			var v float64 = 0.0
			var p int = 0
			for col := 0; col < len(item); col++ {
				if col == classcol {
					continue
				}
				v += root.weightVector[p] * float64(item[col])
				p++
			}
			/* sigmoid */
			v = 1 / (1 + math.Exp(-v))
			/* label 获取 */
			label := root.GetLable(item[classcol])
			error := label - v
			p = 0
			for index, _ := range root.weightVector {
				if p == classcol {
					p++
				}
				root.weightVector[index] += error * alpha * float64(item[p])
				p++
			}
		}
	}

	if len(root.classlabelL) >= 2 {
		root.lchild = LogisticRegressionTrain(DeleteLabelInSet(dataset, classcol, root.classlabelR),
												classcol, alpha, iter)
	}

	if len(root.classlabelR) >= 2 {
		root.rchild = LogisticRegressionTrain(DeleteLabelInSet(dataset, classcol, root.classlabelL),
												classcol, alpha, iter)
	}

	return root
}



/* 逻辑回归 分类算法 */
func LogisticRegression(dataset [][]int, testset [][]int, classcol int, alpha float64, iter int) []int {
	var ret []int = make([]int, 0)
	rootNode := LogisticRegressionTrain(dataset, classcol, alpha, iter)

	for _, item := range testset {
		ret = append(ret, rootNode.LogisticPredict(item, classcol))
	}
	return ret
}