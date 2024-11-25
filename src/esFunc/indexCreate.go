package esFunc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func CreateIndex(indexName string, settings map[string]interface{}) error {
	settingsJSON, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	esClient, err := GetClient()
	if err != nil {
		return fmt.Errorf("es conn is error: %s", err.Error())
	}
	// Create an index request
	req := bytes.NewReader(settingsJSON)

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
