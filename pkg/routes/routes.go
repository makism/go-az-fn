package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type InvokeRequest struct {
	Data     map[string]interface{}
	Metadata map[string]interface{}
}
type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}

func HandleBlobTrigger(w http.ResponseWriter, r *http.Request) {
	var invokeReq InvokeRequest
	d := json.NewDecoder(r.Body)
	decodeErr := d.Decode(&invokeReq)
	if decodeErr != nil {
		// bad JSON or unrecognized json field
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("The JSON data is:invokeReq metadata......")
	fmt.Println(invokeReq.Metadata)

	returnValue := invokeReq.Data["triggerBlob"]
	invokeResponse := InvokeResponse{Logs: []string{"test log1", "test log2"}, ReturnValue: returnValue}

	js, err := json.Marshal(invokeResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

type GreetMessage struct {
	Message string `json:"message"`
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	greet := GreetMessage{
		Message: "hey hi",
	}

	jsonData, err := json.Marshal(greet)

	if r.Method == "GET" {
		if err != nil {
			// handle error
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	} else {
		body, _ := ioutil.ReadAll(r.Body)
		w.Write(body)
	}
}