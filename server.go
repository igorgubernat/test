package main

import (
	"log"
	"net/http"

	"github.com/igorgubernat/test/database"
)

func main() {
    http.HandleFunc("/", database.Handler)
    log.Fatal(http.ListenAndServe(":8888", nil))
}