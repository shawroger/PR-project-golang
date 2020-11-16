package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	// MainDir 图片主目录
	MainDir = "data/png"
)

// FileList 图片文件主目录信息
type FileList struct {
	PathList []string
	Value    int
}

// AnalysisList 图片分析结果
type AnalysisList struct {
	Vector []Vector
	Value  int
}

// Vector 特征向量
type Vector []int

func getFileDir() []string {
	var list []string
	for i := 0; i < 10; i++ {
		str := []string{MainDir, "/", strconv.Itoa(i)}
		list = append(list, strings.Join(str, ""))
	}

	return list
}

// GetFileList 返回各个待分析图片文件
func GetFileList() []FileList {
	var fileList []FileList
	list := getFileDir()
	for value, dir := range list {
		var temp FileList
		temp.Value = value
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			str := []string{dir, "/", file.Name()}
			temp.PathList = append(temp.PathList, strings.Join(str, ""))
		}
		fileList = append(fileList, temp)
	}
	return fileList
}

// RunAnalysis 分析所有文件特征向量
func RunAnalysis() []AnalysisList {
	var vectorList []AnalysisList
	fileList := GetFileList()
	for _, item := range fileList {
		var temp AnalysisList
		temp.Value = item.Value
		for _, file := range item.PathList {
			vector := ImageLoadVector(file)
			temp.Vector = append(temp.Vector, vector)
		}

		vectorList = append(vectorList, temp)
	}

	return vectorList
}
