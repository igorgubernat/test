package parser

import (
	"github.com/igorgubernat/test/model"
)

func parse(line []string) model.Record {
	return model.Record{
		ID:          line[0],
		Name:        line[1],
		Email:       line[2],
		PhoneNumber: line[3],
	}
}
