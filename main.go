package main

import (
	"fmt"
	"log"
)

const (
	fileSource = "./data/test.png"
)

func dealError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func main() {
	analysisList := RunAnalysis()
	vector := ImageLoadVector(fileSource)
	classifyResult := RunClassify(vector, analysisList)
	for _, result := range classifyResult {
		fmt.Println(result)
	}

	fmt.Printf("The result may be %d\n", classifyResult[0].Value)
	fmt.Scanln()
}
