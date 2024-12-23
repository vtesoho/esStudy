package esFunc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

func CreateIndex(indexName string, settingsJSON string) error {

	isExists, err := IndexExists(indexName)
	if err != nil {
		return err
	}
	if isExists {
		return errors.New("index exist")
	}
	var settings map[string]interface{}
	if err := json.Unmarshal([]byte(settingsJSON), &settings); err != nil {
		return fmt.Errorf("error unmarshalling settings JSON: %v", err)
	}

	esClient, err := GetClient()
	if err != nil {
		return fmt.Errorf("es conn is error: %s", err.Error())
	}
	// Create an index request
	req := bytes.NewReader([]byte(settingsJSON))

	// Send the create index request
	res, err := esClient.Indices.Create(indexName, esClient.Indices.Create.WithBody(req))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check the response status
	if res.IsError() {
		return fmt.Errorf("Error creating index: %s", res.String())
	}

	log.Printf("Index %s created successfully", indexName)
	return nil
}
