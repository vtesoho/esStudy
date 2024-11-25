package esFunc

import (
	"bytes"
	"errors"
	"fmt"
)

func DocumentCreate(indexName string, documentId string, documentContent string) error {

	if len(indexName) == 0 {
		return errors.New("indexName is empty")
	}
	if len(documentContent) == 0 {
		return errors.New("documentContent is empty")
	}

	esClient, err := GetClient()
	if err != nil {
		return fmt.Errorf("es conn is error: %s", err.Error())
	}
	// 构建请求
	req := bytes.NewReader([]byte(documentContent))
	res, err := esClient.Index(
		indexName,
		req,
		esClient.Index.WithDocumentID(documentId),
		esClient.Index.WithRefresh("true"),
	)
	if err != nil {
		return fmt.Errorf("failed to index document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing document: %s", res.Status())
	}

	fmt.Printf("Document indexed successfully: %s\n", res.String())
	return nil
}
