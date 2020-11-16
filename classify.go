package main

import "sort"

// ClassifyResult 分类结果项目
type ClassifyResult struct {
	Value    int
	Distance float64
}

// ClassifyList 分类结果
type ClassifyList []ClassifyResult

func calcDistance(vector Vector, sample Vector) float64 {
	var distance float64 = 0
	for index := range vector {
		distance += float64((vector[index] - sample[index]) * (vector[index] - sample[index]))
	}

	return distance
}

// RunClassify 进行分类
func RunClassify(vector Vector, analysisList []AnalysisList) ClassifyList {
	var classifyResult ClassifyList
	for _, unit := range analysisList {
		var distance float64 = 0
		var classifyUnit ClassifyResult
		classifyUnit.Value = unit.Value
		for _, sample := range unit.Vector {
			distance += calcDistance(vector, sample)
		}
		classifyUnit.Distance = distance / float64(len(unit.Vector))
		classifyResult = append(classifyResult, classifyUnit)
	}

	sort.Sort(classifyResult)

	return classifyResult
}

// Less 排序[Sort包用]
func (c ClassifyList) Less(i, j int) bool {
	return c[i].Distance < c[j].Distance
}

// Len 计数[Sort包用]
func (c ClassifyList) Len() int {
	return len(c)
}

// Swap 交换数据[ Sort包用]
func (c ClassifyList) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
