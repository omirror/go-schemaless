package main

import (
	"bytes"

	"encoding/json"

	"fmt"

	"github.com/google/uuid"
	"github.com/rbastic/go-schemaless/examples/apiserver/pkg/api"

	"io/ioutil"
	"net/http"
	"time"
)

type Metadata struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	postURL := "http://localhost:4444/api/put"
	fmt.Println("HTTP JSON POST URL:", postURL)

	var err error

	var metadata Metadata
	metadata.FirstName = "Ryan"
	metadata.LastName = "Bastic"

	var putRequest api.PutRequest
	putRequest.Table = "cell"
	putRequest.RowKey = uuid.New().String()
	putRequest.ColumnKey = "BASE"
	putRequest.RefKey = 1

	var metaBody []byte
	metaBody, err = json.Marshal(metadata)
	if err != nil {
		panic(err)
	}
	putRequest.Body = string(metaBody)

	putRequestMarshal, err := json.Marshal(putRequest)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", postURL, bytes.NewBuffer(putRequestMarshal))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("response Body:", string(body))
}
