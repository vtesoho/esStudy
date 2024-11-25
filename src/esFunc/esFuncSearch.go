package esFunc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

func Search(indexName string, searchQuery map[string]interface{}) (int64, []string, error) {
	var total int64
	var result []string
	if len(indexName) == 0 {
		return total, result, errors.New("indexName is empty")
	}

	esClient, err := GetClient()
	if err != nil {
		return total, result, fmt.Errorf("es conn is error: %s", err.Error())
	}

	// 构建一个搜索请求
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		return total, result, err
	}

	resp, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex(indexName), // 替换为你的索引名称
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)

	if err != nil {
		return total, result, err
	}
	defer resp.Body.Close()

	// 处理响应
	if resp.IsError() {
		return total, result, err
	}

	var res struct {
		Hits struct {
			Total struct {
				Value int64 `json:"value"`
			} `json:"total"`
			Hits []struct {
				ID string `json:"_id"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return total, result, err
	}
	total = res.Hits.Total.Value
	// 提取文档ID
	// var docIDs []string
	for _, hit := range res.Hits.Hits {
		result = append(result, hit.ID)
	}

	return total, result, err

}
