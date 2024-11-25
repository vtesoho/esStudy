package esFunc

import (
	"errors"
	"fmt"
)

func DeleteIndex(indexName string) error {
	if len(indexName) == 0 {
		return errors.New("indexName is empty")
	}
	esClient, err := GetClient()
	if err != nil {
		return fmt.Errorf("es conn is error: %s", err.Error())
	}

	res, err := esClient.Indices.Delete([]string{indexName})
	if err != nil {
		return fmt.Errorf("failed to delete index: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error deleting index: %s", res.Status())
	}

	fmt.Printf("Index '%s' deleted successfully.\n", indexName)
	return nil
}
