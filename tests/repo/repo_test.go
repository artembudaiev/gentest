package repo

import (
	"gentest/repo"
	"testing"
)

func TestGetProduct( t *testing.T){
	fs:=repo.NewFileStorage("../../resources/products.json")
	// #1
	product1,err:=fs.GetProduct(1)
	if err!=nil {
		t.Error("Expected {\n    \"id\" : 1,\n    \"name\": \"Book1\",\n    \"price\": 10\n  }, got ",err.Error())
		return
	}
	expectedProduct:=repo.Product{Id: 1,Name: "Book1",Price: 10}
	if product1 !=expectedProduct {
		t.Error("Expected {\n    \"id\" : 1,\n    \"name\": \"Book1\",\n    \"price\": 10\n  }, got ", product1)
	}
	// #2
	product2,err:=fs.GetProduct(-1)
	if err==nil {
		t.Error("Expected error, got  ",product2)
	}
	if err.Error()!="not found" {
		t.Error("Expected not found error message, got  ", err.Error())
	}
}
