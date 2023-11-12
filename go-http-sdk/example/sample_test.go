package main

import (
	"fmt"
	"net/http"
	"testing"

	sdk "github.com/ibarryyan/go-http-sdk"
)

func TestSDKCreate(t *testing.T) {
	newSDK, err := sdk.NewSDK("http://localhost:9999", "barry", "yan")
	if err != nil && newSDK != nil {
		return
	}
	err1 := newSDK.Create(sdk.CreateRequest{Key: "D", Val: "1"})
	if err1.Code != http.StatusOK {
		fmt.Println(err1)
	}
}

func TestSDKGet(t *testing.T) {
	newSDK, err := sdk.NewSDK("http://localhost:9999", "barry", "yan")
	if err != nil && newSDK != nil {
		return
	}
	resp, err2 := newSDK.Get("D")
	if err2.Code != http.StatusOK {
		fmt.Println(err2)
	}
	fmt.Println(resp)
}
