package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

    "github.com/igorgubernat/test/model"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "could not read record: %v", err)
		fmt.Println(err)
		return
	}

	var record model.Record

	err = json.Unmarshal(body, &record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "could not decode record: %v", err)
		fmt.Println(err)
		return
	}

	save(record)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "saved successfully")
	fmt.Println("success")
}
