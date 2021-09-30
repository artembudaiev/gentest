package main

import (
	"gentest/middlewares"
	"gentest/router"
	"gentest/tests/moke"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	router := router.New()
	mux.Handle("/pay", middlewares.HostnameHeader(http.HandlerFunc(router.PayRoute)))
	mux.HandleFunc("/api/",moke.ApiRoute)
	jsonHeader:=middlewares.JSONHeader(mux)
	http.ListenAndServe(":8080", jsonHeader)
}
