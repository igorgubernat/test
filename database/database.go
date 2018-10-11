package database

import (
	"github.com/igorgubernat/test/model"
)

var db map[string]model.Record

func init() {
	db = make(map[string]model.Record)
}

func save(record model.Record) {
	db[record.ID] = record
}