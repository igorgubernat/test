package parser

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func Read() error {
	file, err := os.Open("data.csv")
	if err != nil {
		return err
	}

	reader := csv.NewReader(bufio.NewReader(file))

	// omit header
	_, err = reader.Read()
	if err != nil {
		return err
	}

	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		record := parse(line)
		err = send(record)
		if err != nil {
			fmt.Printf("error sending record: %v\n", err)
		}
	}

	return nil
}
