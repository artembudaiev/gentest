package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// PayRoute /pay/ basic route
func (r *router) PayRoute(res http.ResponseWriter, req *http.Request) {
	start:=time.Now()

	if req.URL.Path != "/pay" {
		http.Error(res, fmt.Sprintf("Not found %v" ,req.URL.Path), http.StatusNotFound)
		return
	}
	if req.Method == http.MethodGet {
		r.getPayIdRoute(res, req,start)
	} else {
		http.Error(res, fmt.Sprintf("expect method GET /pay/, got %v", req.Method), http.StatusMethodNotAllowed)
		return
	}
}

//getPayIdRoute get /pay route (expects /pay?id=*)
func (r *router) getPayIdRoute(res http.ResponseWriter, req *http.Request, start time.Time) {

	//expect ID parameter
	param,err :=getParam(req.URL.RawQuery,"id")
	if err !=nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	//try to convert ID to int
	id, err := strconv.Atoi(param)
	if err != nil {
		http.Error(res, fmt.Sprintf("expected type int as id, got %T",param), http.StatusBadRequest)
		return
	}

	//get product with id from storage
	product, err := r.storage.GetProduct(id)
	if err != nil {
		res.Header().Set("X-Response-Time",strconv.FormatInt(time.Since(start).Milliseconds(),10))
		res.Write([]byte(`{ "link" : "https://play.google.com/store/apps/details?id=com.headway.books&hl=ru&gl=US" }`))
		return
	}

	//request link from google api
	googleData,err:=apiRequest("http://localhost:8080/api/link/google?price="+strconv.Itoa(product.Price))
	if err != nil {
		res.Header().Set("X-Response-Time",strconv.FormatInt(time.Since(start).Milliseconds(),10))
		log.Println(err.Error())
		res.Write([]byte(`{ "link" : "https://play.google.com/store/apps/details?id=com.headway.books&hl=ru&gl=US" }`))
		return
	}

	//request link from apple api
	appleData,err:=apiRequest("http://localhost:8080/api/link/apple?price="+strconv.Itoa(product.Price))
	if err != nil {
		res.Header().Set("X-Response-Time",strconv.FormatInt(time.Since(start).Milliseconds(),10))
		log.Println(err.Error())
		res.Write([]byte(`{ "link" : "https://play.google.com/store/apps/details?id=com.headway.books&hl=ru&gl=US" }`))
		return
	}


	//response
	res.Header().Set("X-Response-Time",strconv.FormatInt(time.Since(start).Milliseconds(),10))
	res.Write([]byte(fmt.Sprintf(`{ "google" : %v, "apple" : %v }`,string(googleData),string(appleData))))
}

// getParam from raw query executes value by key
func getParam(query, key string) (string,error) {
	m, err := url.ParseQuery(query)
	if err != nil {
		return "",errors.New(fmt.Sprintf("failed to parse %v",query))
	}
	keys, ok := m[key]
	if !ok {
		return "",errors.New(fmt.Sprintf("expect %v parameter, got %v",key,m))
	}
	return keys[0],nil
}

// apiRequest makes GET request to query URL, returns byte data if success
func apiRequest(query string) ([]byte,error) {
	resp, err := http.Get(query)
	if err != nil {
		log.Println(err.Error())
		return nil,errors.New(fmt.Sprintf("failed to send request to %v",query))
	}
	defer resp.Body.Close()
	if resp.StatusCode>299 {
		log.Println(resp.StatusCode)
		return nil,errors.New(fmt.Sprintf("got status code %v",resp.StatusCode))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	if !json.Valid(body){
		return nil,errors.New(fmt.Sprintf("answer %v is not json",body))
	}
	return body,nil
}