package esFunc

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	ESClient *elasticsearch.Client
)

func GetClient() (*elasticsearch.Client, error) {
	if ESClient == nil {
		err := Init()
		if err == nil {
			return ESClient, err
		}
		return ESClient, errors.New("ESClient is not Init")
	}

	return ESClient, nil
}

func Init() error {
	if ESClient != nil {
		return nil
	}
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://10.192.10.11:9200", // 你的 Elasticsearch 地址
		},
		Username: "elastic",
		Password: "Iceb4NUov86fCPeDNtPF",
		// 其他配置，可以添加认证等
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			MaxIdleConns:          200,
			IdleConnTimeout:       time.Second,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	ESClient, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return err
	}
	fmt.Println("ESClient", ESClient)
	return nil

}
