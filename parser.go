package main

import (
	"log"

	"github.com/igorgubernat/test/parser"
)

func main() {
	err := parser.Read()
	if err != nil {
		log.Fatal(err)
	}
}
