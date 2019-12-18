package classific

import (
	"fmt"
	"log"
	"math"
)

/*
决策树分类算法:
@input
	data - 输入数据矩阵
	test - 测试数据集合
	classcol - 预测类别所在列
@output
	[]int - 分类值
*/

type DTNode struct {
	end bool			// 是否为终结点
	class int			// 终结点情况下属于哪一类
	col   int 			// 划分列
	split int			// 划分点 <= split > split 两类
	lchild *DTNode
	rchild *DTNode
}

func (d *DTNode) output() {

	if d == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(d)
	fmt.Print("lchild : ")
	d.lchild.output()
	fmt.Print("rchild : ")
	d.rchild.output()
}

func (d *DTNode) find(testcase []int) int {
	var loop *DTNode = d
	for {
		if loop == nil {
			log.Fatal("error")
			return 0
		}
		if loop.end {
			return loop.class
		}
		v := testcase[loop.col]
		if v <= loop.split {
			if loop.lchild != nil {
				loop = loop.lchild
			}else {
				loop = loop.rchild
			}
		}else {
			if loop.rchild != nil {
				loop = loop.rchild
			}else {
				loop = loop.lchild
			}
		}
	}
}

/*
判断数据某一列是否只有一个值
@input
	data - 输入数据
	col  - 某一列
@output
	bool - 是否只有一个类
	int  - 类名
*/
func IsOneClass(data [][]int, col int) (bool, int) {
	for i := 1; i < len(data); i++ {
		if data[i][col] != data[i-1][col] {
			return false, -1
		}
	}
	return true, data[0][col]
}

/*
判断数据中最多的类
@intput
	data - 输入数据
	col  - 某一列
@output
	int  - 某一列数据中最多的值
*/
func TheMostClass(data [][]int, col int) int {
	var cntMap map[int]int = make(map[int]int)
	for _, item := range data {
		cntMap[item[col]]++
	}

	cnt := math.MinInt64
	ret := 0

	for k, v := range cntMap {
		if v >= cnt {
			ret = k
			cnt = v
		}
	}
	return ret
}

/*
计算熵
@input
	data     - 输入矩阵
	classcol - 用于计算信息熵的列
*/
func ColEntropy(data [][]int, classcol int) float64 {
	classpro := Probability(data, classcol)
	classentropy := 0.0
	for _, v := range classpro {
		entropy := -(v * math.Log2(v))
		classentropy += entropy
	}
	return classentropy
}

/*
按某一列的某值划分
*/
func SplitByValueAndCol(dataset [][]int, col, value int, colSplitType map[int]bool) [2][][]int {
	var ret [2][][]int
	t := colSplitType[col]

	for _, item := range dataset {
		if t {
			if item[col] <= value {
				ret[0] = append(ret[0], item)
			}else {
				ret[1] = append(ret[1], item)
			}
		}else {
			if item[col] == value {
				ret[0] = append(ret[0], item)
			}else {
				ret[1] = append(ret[1], item)
			}
		}
	}

	return ret
}

/*
找到划分点
@input
	dataset  - 数据集
	col      - 特征列
@output

*/
func FindSliptPoint(dataset [][]int, col, classcol int, colSplitType map[int]bool) int{
	var tosplitpoint map[int][]int = make(map[int][]int)
	var entropytable map[int]float64 = make(map[int]float64)
	var minK int = 0
	var minV float64 = math.MaxFloat64

	for _, item := range dataset {
		tosplitpoint[item[col]] = append(tosplitpoint[item[col]], item[classcol])
	}

	for k, _ := range tosplitpoint {
		diffset := SplitByValueAndCol(dataset, col, k, colSplitType)
		entropytable[k] = ColEntropy(diffset[0], classcol) + ColEntropy(diffset[1], classcol)
	}

	for k, v := range entropytable {
		if v <= minV {
			minV = v
			minK = k
		}
	}

	return minK
}

/*
创建决策树
@input
	dataset  - 数据集
	classcol - 按某列分类
	mask     - 标记哪些列已被使用
@output

*/
func BuildDecisionTree(dataset [][]int, classcol int, mask map[int]bool, colSplitType map[int]bool) *DTNode{
	// 集合为空
	if len(dataset) == 0 {
		return nil
	}
	col := len(dataset[0])
	// 属性集
	oneclass, classlabel := IsOneClass(dataset, classcol)
	if oneclass {
		return &DTNode{
			end:    true,
			class:  classlabel,
			col:    0,
			split:  0,
			lchild: nil,
			rchild: nil,
		}
	}
	// 集合为空
	if len(mask) == col {
		return &DTNode{
			end:    true,
			class:  TheMostClass(dataset, classcol),
			col:    0,
			split:  0,
			lchild: nil,
			rchild: nil,
		}
	}
	// 初始信息熵
	shannolentropy := ColEntropy(dataset, classcol)
	// 求某一列的条件熵
	var increentropy map[int]float64 = make(map[int]float64)
	var maxIndex = -math.MaxFloat64
	var maxPos = 0
	for i := 0; i < col; i++ {
		if i == classcol {
			continue
		}
		if _, ok := mask[i]; ok {
			continue
		}

		vectorCol := SplitByCol(dataset, i)
		vectorPro := Probability(dataset, i)

		vectorValue := 0.0
		for k, v := range vectorCol {
			vectorValue += ColEntropy(v, classcol) * vectorPro[k]
		}
		increentropy[i] = shannolentropy - vectorValue
		if increentropy[i] > maxIndex {
			maxIndex = increentropy[i]
			maxPos = i
		}
	}

	mask[maxPos] = true

	splitV := FindSliptPoint(dataset, maxPos, classcol, colSplitType)
	diffset := SplitByValueAndCol(dataset, maxPos, splitV, colSplitType)

	var rootNode DTNode = DTNode{
		end:    false,
		class:  0,
		col:    maxPos,
		split:  splitV,
		lchild: nil,
		rchild: nil,
	}

	rootNode.lchild = BuildDecisionTree(diffset[0], classcol, mask, colSplitType)
	rootNode.rchild = BuildDecisionTree(diffset[1], classcol, mask, colSplitType)

	return &rootNode
}

func DecisionTree(dataset [][]int, testset [][]int, classcol int, colSplitType map[int]bool) []int{
	var ret []int
	var mask map[int]bool = make(map[int]bool)
	mask[classcol] = true
	root := BuildDecisionTree(dataset, classcol, mask, colSplitType)

	for _, item := range testset {
		r := root.find(item)
		ret = append(ret, r)
	}
	return ret
}
