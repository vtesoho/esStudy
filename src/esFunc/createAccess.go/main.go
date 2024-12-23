package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	esURL       = "http://172.28.251.218:9200"
	username    = "elastic"
	password    = "Iceb4NUov86fCPeDNtPF" // 请替换为您的密码
	serviceName = "kibana_service_account"
)

type CreateServiceAccountRequest struct {
	Name string `json:"name"`
}

type GenerateTokenRequest struct {
	Name    string `json:"name"`
	Service string `json:"service"`
}

func main() {
	// 创建服务账户
	if err := createServiceAccount(); err != nil {
		fmt.Printf("Error creating service account: %v\n", err)
		return
	}

	// 生成服务账户令牌
	token, err := generateServiceAccountToken()
	if err != nil {
		fmt.Printf("Error generating token: %v\n", err)
		return
	}

	fmt.Printf("Generated token: %s\n", token)
}

func createServiceAccount() error {
	reqBody := CreateServiceAccountRequest{Name: serviceName}
	reqBodyJSON, _ := json.Marshal(reqBody)

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/_security/service/credentials", esURL), bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		return err
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create service account: %s", body)
	}

	return nil
}

func generateServiceAccountToken() (string, error) {
	reqBody := GenerateTokenRequest{Name: "kibana-token", Service: serviceName}
	reqBodyJSON, _ := json.Marshal(reqBody)

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/_security/service/token", esURL), bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to generate token: %s", body)
	}

	var responseMap map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&responseMap); err != nil {
		return "", err
	}

	return responseMap["token"], nil
}
