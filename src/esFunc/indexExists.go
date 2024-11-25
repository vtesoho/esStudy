package esFunc

import (
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func IndexExists(indexName string) (bool, error) {
	ctx := context.Background()
	esClient, err := GetClient()
	if err != nil {
		return false, fmt.Errorf("es conn is error: %s", err.Error())
	}

	req := esapi.IndicesExistsRequest{
		Index: []string{indexName},
	}

	res, err := req.Do(ctx, esClient)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.IsError() {
		if res.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error: %s", res.String())
	}

	return true, nil
}
