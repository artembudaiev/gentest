package router

import (
	"gentest/repo"
)

type router struct {
	storage repo.Storage
}

func New() *router {
	// if ypu want to store smth to db, implement Storage interface and change code here
	return &router{storage: repo.NewFileStorage("resources/products.json")}
}
