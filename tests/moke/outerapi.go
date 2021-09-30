package moke

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func ApiRoute(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, fmt.Sprintf("expect method GET, got %v", req.Method), http.StatusMethodNotAllowed)
		return
	}
	//expect Price parameter
	param, err := getParam(req.URL.RawQuery, "price")
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	//try to convert price to int
	price, err := strconv.Atoi(param)
	if err != nil {
		http.Error(res, fmt.Sprintf("expected type int as price, got %T", param), http.StatusBadRequest)
		return
	}
	if req.URL.Path == "/api/link/google" {
		res.Write([]byte(fmt.Sprintf(`{"link":"/api/form/google?price=%v" }`,price)))
	} else if req.URL.Path == "/api/link/apple" {
		res.Write([]byte(fmt.Sprintf(`{"link":"/api/form/apple?price=%v" }`,price)))
	} else if req.URL.Path == "/api/form/google" {
		res.Write([]byte(fmt.Sprintf(`{"form":"This is GooglePay form! Pay %v"}`, price)))
	} else if req.URL.Path == "/api/form/apple" {
		res.Write([]byte(fmt.Sprintf(`{"form":"This is ApplePay form! Pay %v"}`, price)))
	} else {
		http.Error(res, fmt.Sprintf("unexpected path %v", req.URL.Path), http.StatusMethodNotAllowed)
		return
	}
}
func getParam(query, key string) (string, error) {
	m, err := url.ParseQuery(query)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to parse %v", query))
	}
	keys, ok := m[key]
	if !ok {
		return "", errors.New(fmt.Sprintf("expect %v parameter, got %v", key, m))
	}
	return keys[0], nil
}
