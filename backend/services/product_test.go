package services

import (
	"testing"
)

func Test_CreateProduct(t *testing.T) {
	data := ProductService{
		Items: []Product{{
			Id:         "3",
			IdCategory: "1",
			Name:       "Quần nam 3",
		}},
	}
	err := data.Create()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateProduct(t *testing.T) {
	data := ProductService{
		Items: []Product{{
			Id:         "3",
			IdCategory: "1",
			Name:       "Quần dài",
		}},
	}
	err := data.Update()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetProduct(t *testing.T) {
	data := ProductService{Items: []Product{{
		Id: "3",
	}}}
	err := data.Get()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetAllProduct(t *testing.T) {
	limit := 10
	page := 1
	data := ProductService{}
	err := data.GetAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteProduct(t *testing.T) {
	data := ProductService{Items: []Product{{
		Id: "3",
	}}}
	err := data.Delete()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
