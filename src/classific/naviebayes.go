package classific

/*
贝叶斯分类算法:
@input
	data - 输入数据矩阵
	test - 测试数据集合
	classcol - 预测类别所在列
@output
	[]int - 分类值
*/
func NavieBayes(dataset [][]int, testset [][]int, classcol int) []int{
	var classVector map[int][][]int = SplitByCol(dataset, classcol) // 分类矩阵map[类别名] [][]int 矩阵值
	var classpro map[int]float64 = Probability(dataset, classcol)	// 类别概率map[类别名] 概率
	var classname []int = make([]int, 0)							// 类别名

	// 条件概率 map[类别名] map[第k参数] map[第k参数值]概率
	var classcondpro map[int]map[int]map[int]float64 = make(map[int]map[int]map[int]float64)
	// 全部概率 map[第k参数]map[第k参数值]概率
	var allpro map[int]map[int]float64 = make(map[int]map[int]float64)

	col := len(dataset[0])

	for k, v := range classVector {
		classname = append(classname, k)
		for i := 0; i < col; i++ {
			if i == classcol {
				continue
			}
			if _, ok := classcondpro[k]; !ok {
				classcondpro[k] = make(map[int]map[int]float64)
			}
			classcondpro[k][i] = Probability(v, i)
			allpro[i] = Probability(dataset, i)
		}
	}

	var pre []int = make([]int, 0)
	for _, item := range testset {
		var maxIndex int = 0
		var maxPro float64 = 0.0

		for _, c := range classname {
			var fz1 float64 = 1.0
			var fz2 float64 = 1.0
			var fm float64 = 1.0

			fz2 = classpro[c]

			for col, v := range item {
				if col == classcol {
					continue
				}
				if allpro[col][v] == 0 {
					fm *= float64(1) / float64(len(dataset) + 1)
				}else {
					fm *= allpro[col][v]
				}
				if classcondpro[c][col][v] == 0 {
					fz1 *= float64(1) / float64(len(classVector[c]) + 1)
				}else {
					fz1 *= classcondpro[c][col][v]
				}
			}
			if (fz1 * fz2 / fm) > maxPro {
				maxIndex = c
				maxPro = fz1 * fz2 / fm
			}
		}
		pre = append(pre, maxIndex)
	}
	return pre
}

