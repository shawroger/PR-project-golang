package main

import (
	"log"
)

func dealError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {}
