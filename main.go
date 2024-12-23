package main

import (
	"encoding/json"
	"esStudy/src/conn/viVoiceCheckLogsMysql"
	"esStudy/src/esFunc"
	"esStudy/src/gorms/aiVoice/clientManageLogsGorm"
	"fmt"
	"log"
	"strconv"
)

func main() {
	// CreateIndexText()
	IndexExistsTest()

}

func CreateIndexText() {
	// JSON configuration for the index
	settingsJSON := `{
		"settings": {
			"number_of_shards": 3,
			"number_of_replicas": 0,
			"index": {
				"refresh_interval": "1m"
			}
		},
		"mappings": {
			"properties": {
				"id": {
					"type": "integer"
				},
				"agentname": {
					"type": "keyword"
				},
				"agentaccount": {
					"type": "keyword"
				},
				"callere164": {
					"type": "keyword"
				},
				"calleraccesse164": {
					"type": "keyword"
				},
				"calleee164": {
					"type": "keyword"
				},
				"calleeaccesse164": {
					"type": "keyword"
				},
				"callerip": {
					"type": "keyword"
				},
				"callerrtpip": {
					"type": "keyword"
				},
				"callergatewayid": {
					"type": "keyword"
				},
				"starttime": {
					"type": "long"
				},
				"stoptime": {
					"type": "long"
				},
				"holdtime": {
					"type": "integer"
				}
			}
		}
	}`

	indexName := "20241210_73_11_12"

	if err := esFunc.CreateIndex(indexName, settingsJSON); err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

}

type RecordInfo struct {
	IsFinal bool   `json:"is_final"`
	Mode    string `json:"mode"`
	Text    string `json:"text"`
	WavName string `json:"wav_name"`
}

func InsertTest() {
	indexName := "my_index"
	tableName := "20241019logs_74_831m2"
	db := viVoiceCheckLogsMysql.GetMysqlDB()
	var clientManageLogsInfos []clientManageLogsGorm.Table
	db.Table(tableName).Where("Id < 10").Find(&clientManageLogsInfos)

	for _, v := range clientManageLogsInfos {
		var recordInfo RecordInfo
		json.Unmarshal([]byte(v.RightResult), &recordInfo)
		var insert CheckTest
		idStr := strconv.Itoa(v.Id)
		insert.Id = v.Id
		insert.Content = recordInfo.Text

		insertStr, _ := json.Marshal(insert)
		esFunc.DocumentCreate(indexName, idStr, string(insertStr))
	}

}

func SearchTest() {
	indexName := "20241125_74_231"
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"Content": "加微信",
			},
		},
		"from": "0",
		"size": "10",
	}
	total, results, err := esFunc.Search(indexName, searchQuery)

	fmt.Println("total, results, err", total, results, err)
}

func IndexExistsTest() {
	indexName := "20241210_73_11_12"
	// indexName := "my_index"
	result, err := esFunc.IndexExists(indexName)
	fmt.Println("result,err", result, err)
}

type CheckTest struct {
	Id      int    `json:"Id"`
	Content string `json:"Content"`
}
