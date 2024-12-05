package main

import (
	"encoding/json"
	"net/http"
)

type Params struct {
	Name string
	Age  int
	Sex  string
}

type Response struct {
	Success bool
	Message string
}

func createPersonalData(params *Params) *Response {
	return nil
}

func createPersonalDataHandler(rw http.ResponseWriter, req *http.Request) {

	var params Params
	err := json.NewDecoder(req.Body).Decode(&params)
	if err != nil {
		http.Error(rw, "Invalid Input", http.StatusBadRequest)
	}

}

func main() {
	http.HandleFunc("/personal_data/create", createPersonalDataHandler)
	http.ListenAndServe(":8080", nil)
}
