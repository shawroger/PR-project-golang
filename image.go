package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
)

// ColorList 像素颜色数组
// 长度 * 宽度 * [r, b, g, a]
type ColorList [][][4]uint32

// ImageInfo 图片信息
// colorList 颜色序列
// width 图片宽度
// width 图片长度
type ImageInfo struct {
	Color  ColorList
	Width  int
	Height int
}

// ImageLoad 读取并图片信息
// colorList 颜色序列
// width 图片宽度
// width 图片长度
func ImageLoad(filePath string) ImageInfo {

	// 打开图片文件
	file, err := os.Open(filePath)

	//处理异常，下同 [At main.go]
	dealError(err)

	// 异步关闭文件
	defer file.Close()

	// 解析图片文件
	pixel, _, err := image.Decode(file)
	dealError(err)

	// 获取图片大小
	pixelSize := pixel.Bounds().Max

	// 最终的像素RGBA数组
	var colorList ColorList

	for y := 0; y < pixelSize.Y; y++ {
		var colorRow [][4]uint32
		for x := 0; x < pixelSize.X; x++ {
			r, g, b, a := pixel.At(x, y).RGBA()
			colorPoint := [4]uint32{r, g, b, a}
			colorRow = append(colorRow, colorPoint)
		}
		colorList = append(colorList, colorRow)
	}

	return ImageInfo{colorList, pixelSize.X, pixelSize.Y}
}

func isBlack(colorPoint [4]uint32) bool {
	return (colorPoint[0] == 0 && colorPoint[1] == 0 && colorPoint[2] == 0 && colorPoint[3] == 65535)
}

// RasterizeList 像素颜色数组
// 长度 * 宽度 * [0|1]
type RasterizeList [][]byte

// ImageRasterize 返回图片点阵数组
func ImageRasterize(imgInfo ImageInfo) RasterizeList {
	// 定义点阵数组
	var rasterizeList RasterizeList
	for _, row := range imgInfo.Color {
		var rasterizeRow []byte

		for _, color := range row {
			// 判断黑色，是1否0
			var flag byte = 0

			if isBlack(color) {
				flag = 1
			}
			rasterizeRow = append(rasterizeRow, flag)
		}

		rasterizeList = append(rasterizeList, rasterizeRow)
	}

	return rasterizeList
}

func vectorize(vectorGroup [][]RasterizeList) []int {
	var vector []int
	for i := 0; i < len(vectorGroup); i++ {
		for j := 0; j < len(vectorGroup[0]); j++ {
			counter := 0
			chunk := vectorGroup[i][j]
			for ii := 0; ii < len(chunk); ii++ {
				for jj := 0; jj < len(chunk); jj++ {
					if chunk[ii][jj] == 1 {
						counter++
					}
				}
			}
			vector = append(vector, counter)
		}
	}

	return vector
}

// VectorRasterize 点阵向量化
func VectorRasterize(imgInfo ImageInfo) []int {
	var (
		vectorGroup     [][]RasterizeList
		vectorGroupItem []RasterizeList
	)
	rasterizeList := ImageRasterize(imgInfo)
	offset := int(math.Floor(float64(imgInfo.Width / 4)))

	for i := 0; i < len(rasterizeList); i = i + offset {
		vectorGroupItem = append(vectorGroupItem, rasterizeList[i:i+offset])
	}

	for i := 0; i < len(rasterizeList); i = i + offset {
		var temp []RasterizeList
		for _, group := range vectorGroupItem {
			var temp2 RasterizeList
			for _, item := range group {
				temp2 = append(temp2, item[i:i+offset])
			}
			temp = append(temp, temp2)
		}

		vectorGroup = append(vectorGroup, temp)
	}

	return vectorize(vectorGroup)
}

// ImageLoadVector 直接返回图片特征向量
func ImageLoadVector(filePath string) []int {
	imageInfo := ImageLoad(filePath)
	vector := VectorRasterize(imageInfo)
	return vector
}

// ImageLoadRasterize 直接返回点阵数据
func ImageLoadRasterize(filePath string) RasterizeList {
	imageInfo := ImageLoad(filePath)
	rasterize := ImageRasterize(imageInfo)
	return rasterize
}
