package services

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_GetProduct(t *testing.T) {
	// Create service and assign to data
	data := ProductService{Items: []models.Product{{
		Id: "3",
	}}}

	// Execute method and if error happen send error
	err := data.Get()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetAllProduct(t *testing.T) {
	// Create limit, page, service and assign to data
	limit := 10
	page := 1
	data := ProductService{}

	// Execute method and if error happen send error
	err := data.GetAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_CreateProduct(t *testing.T) {
	// Create service and assign to data
	data := ProductService{
		Items: []models.Product{{
			Id:         "3",
			IdCategory: "1",
			Name:       "Quần nam 3",
		}},
	}

	// Execute method and if error happen send error
	err := data.Create()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateProduct(t *testing.T) {
	// Create service and assign to data
	data := ProductService{
		Items: []models.Product{{
			Id:         "3",
			IdCategory: "1",
			Name:       "Quần dài",
		}},
	}

	// Execute method and if error happen send error
	err := data.Update()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteProduct(t *testing.T) {
	// Create service and assign to data
	data := ProductService{Items: []models.Product{{
		Id: "3",
	}}}

	// Execute method and if error happen send error
	err := data.Delete()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
