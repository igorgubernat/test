package main

import (
	"github.com/igorgubernat/test/parser"
	"log"
)

func main() {
	err := parser.Read()
	if err != nil {
		log.Fatal(err)
	}
}
