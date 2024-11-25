package main

import (
	"encoding/json"
	"esStudy/src/conn/viVoiceCheckLogsMysql"
	"esStudy/src/esFunc"
	"esStudy/src/gorms/aiVoice/clientManageLogsGorm"
	"fmt"
	"strconv"
)

func main() {
	// defaultES.Init()

	// viVoiceCheckLogsMysql.InitMysql()

	// settings := map[string]interface{}{
	// 	"settings": map[string]interface{}{
	// 		"analysis": map[string]interface{}{
	// 			"tokenizer": map[string]interface{}{
	// 				"ik_smart": map[string]interface{}{
	// 					"type": "ik_smart",
	// 				},
	// 			},
	// 			"analyzer": map[string]interface{}{
	// 				"ik_smart_analyzer": map[string]interface{}{
	// 					"type":      "custom",
	// 					"tokenizer": "ik_smart",
	// 				},
	// 			},
	// 		},
	// 	},
	// 	"mappings": map[string]interface{}{
	// 		"properties": map[string]interface{}{
	// 			"Id": map[string]interface{}{
	// 				"type": "integer",
	// 			},
	// 			"Content": map[string]interface{}{
	// 				"type":     "text",
	// 				"analyzer": "ik_smart_analyzer",
	// 			},
	// 		},
	// 	},
	// }

	// err := esFuncCreateIndex.CreateIndex(indexName, settings)
	// if err != nil {
	// 	fmt.Printf("Error creating index: %s\n", err)
	// }
	// fmt.Println("esStudy start")

	// esFuncDeleteIndex.DeleteIndex(indexName)

	SearchTest()
	// IndexExistsTest()

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
	indexName := "my_index"
	result, err := esFunc.IndexExists(indexName)
	fmt.Println("result,err", result, err)
}

type CheckTest struct {
	Id      int    `json:"Id"`
	Content string `json:"Content"`
}
