package classific

/*
	计算F1值
@input
	testset - 测试(真实验证)值 preset - 预测值
	class - 分类数量
@output
	float64 - F1值
	float64 - 正确率
	float64 - 召回率
*/
func F1(preset, testset []int) (float64, float64, float64){
	var sample map[int]int = make(map[int]int)/* 测试值(真实验证)中每类的数量 */
	var pre map[int]int = make(map[int]int)   /* 预测值(分类预测)中每类的数量 */
	var ans map[int]int = make(map[int]int)   /* 测试出准确的每类数量 */

	var predic map[int]float64 = make(map[int]float64)	/* 正确率 */
	var recall map[int]float64 = make(map[int]float64)	/* 召回率 */

	classnum := 0

	for i := 0; i < len(preset); i++ {
		sample[testset[i]]++
		pre[preset[i]]++
		if preset[i] == testset[i] {
			ans[preset[i]]++
		}
	}

	mpredic := 0.0	// 宏正确率
	mrecall := 0.0  // 宏召回率
	for k, _ := range sample {
		predic[k] = float64(ans[k]) / float64(pre[k])
		recall[k] = float64(ans[k]) / float64(sample[k])
		mpredic += predic[k]
		mrecall += recall[k]
		classnum++
	}
	mpredic /= float64(classnum)
	mrecall /= float64(classnum)
	return 2 * mpredic * mrecall / (mpredic + mrecall), mpredic, mrecall
}

/* 求解第col列元素的概率 */
func Probability(data [][]int, col int) map[int]float64{
	var ret map[int]float64 = make(map[int]float64)
	for row, _ := range data {
		ret[data[row][col]]++
	}
	for k, v := range ret {
		ret[k] = v / float64(len(data))
	}
	return ret
}

