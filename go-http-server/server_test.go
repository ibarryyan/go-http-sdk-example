package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestCreateAPI(t *testing.T) {
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 创建POST请求的body
	reqData := []byte(`{"key":"A","val":"1"}`)

	// 创建一个POST请求
	req, err := http.NewRequest("POST", "http://localhost:9999/create", bytes.NewBuffer(reqData))
	if err != nil {
		fmt.Println("创建请求时发生错误:", err)
		return
	}

	// 设置请求头，如果有特殊需求的话
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("name", "barry yan")

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求时发生错误:", err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应时发生错误:", err)
		return
	}

	// 打印响应内容
	fmt.Println(string(body))
}

func TestGetAPI(t *testing.T) {
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 创建一个GET请求
	req, err := http.NewRequest("GET", "http://localhost:9999/get/A", nil)
	if err != nil {
		fmt.Println("创建请求时发生错误:", err)
		return
	}
	req.Header.Set("name", "barry yan")

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求时发生错误:", err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// 读取响应的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应时发生错误:", err)
		return
	}

	// 打印响应内容
	fmt.Println(string(body))
}
