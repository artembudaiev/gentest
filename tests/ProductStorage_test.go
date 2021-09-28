package product

import (
	product "gentest/models"
	"testing"
)

func TestGetProduct(t *testing.T){
	fs:=product.NewFileStorage()
	if product, err :=fs.GetProduct(1); err!=nil {
		t.Error(err)
	} else if product.Id!=1{
		t.Error("Got :",product," while expecting id = ",1)
	}
}