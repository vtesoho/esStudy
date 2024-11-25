package defaultES

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	ESClient *elasticsearch.Client
)

func Init() {
	if ESClient != nil {
		return
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
	}

}
