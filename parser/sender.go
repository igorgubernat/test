package parser

import (
	"net/http"
	"encoding/json"
	"bytes"
	"time"
	"fmt"
	"github.com/igorgubernat/test/model"
)

const dbUrl = "http://localhost:8888"
const httpRequestTimeout = 5 * time.Second

func send(record model.Record) error {
	reqBody, err := json.Marshal(record)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, dbUrl, bytes.NewReader(reqBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := http.DefaultClient
	client.Timeout = httpRequestTimeout
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("could not create record: database service returned status %d", res.StatusCode)
	}

	return nil
}