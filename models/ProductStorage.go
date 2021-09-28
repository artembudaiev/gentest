package product

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Storage interface {
	GetProduct(id int) (Product, error)
}

type FileStorage struct {
	path     string
}

func NewFileStorage() *FileStorage {
	return &FileStorage{
		path: "../resources/products.json",
	}
}

func (pfs *FileStorage) readData() ([]Product,error) {
	var (
		dataFromFile []byte
		error        error
		products []Product
	)

	if dataFromFile, error = ioutil.ReadFile(pfs.path); error != nil {
		return nil,error
	}

	if error = json.Unmarshal(dataFromFile, &products); error != nil {
		return nil,error
	}
	return products,nil
}

func (pfs *FileStorage) GetProduct(id int) (Product, error) {
	var (
		error error
		products []Product
	)
	if products,error = pfs.readData(); error !=nil {
		return Product{}, error
	}
	for _,product := range products {
		if product.Id==id {
			return product, nil
		}
	}
	return Product{}, errors.New("not found")
}

//type DBStorage struct {}
//
//func (D DBStorage) GetProduct(id int) (Product, error) {
//	panic("implement me")
//}
