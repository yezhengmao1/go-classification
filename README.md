# 简介
使用`golang`实现三种分类算法：
* 朴素贝叶斯（Naive Bayes）
* 逻辑回归（Logistic Regression）
* 决策树（Decision Tree）

# 文件说明
* data/ - 来自于UCI数据集文件
* src/data/data.go - 用于将数据集转化为矩阵
* src/classific/decisiontree.go - 决策树算法
* src/classific/naviebayes.go - 朴素贝叶斯算法
* src/classific/logisticregression.go - 逻辑回归算法
* src/classific/classific.go - 指标测试函数F1值,召回率,正确率
* src/main.go - 主函数

# 主函数 - 参数说明
* setnum - K折交叉验证
* classcol - 由data.go转换数据为矩阵，类别所在列号
* alldata - 由data.go转化数据为矩阵

# 测试指标 - 以occupancy.txt数据为例子
朴素贝叶斯: F1:0.9667 召回率:0.9743 正确率:0.9593   
决策树   :F1:0.9822 召回率:0.9909 正确率:0.9737   
逻辑回归:  F1:0.8806 召回率:0.8971 正确率:0.8696   

 
