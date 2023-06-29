package services

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_CreateCategory(t *testing.T) {
	// Create service and assign to data
	data := CategoryService{
		Items: []models.Category{{Id: "3", Name: "Nón"}},
	}

	// Execute method and if error happen send error
	err := data.Create()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteCategory(t *testing.T) {
	// Create service and assign to data
	data := CategoryService{Items: []models.Category{{
		Id: "3",
	}}}

	// Execute method and if error happen send error
	err := data.Delete()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetCategory(t *testing.T) {
	// Create service and assign to data
	data := CategoryService{Items: []models.Category{{
		Id: "1",
	}}}

	// Execute method and if error happen send error
	err := data.Get()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetAllCategory(t *testing.T) {
	// Create limit, page, service and assign to data
	limit := 10
	page := 1
	data := CategoryService{}

	// Execute method and if error happen send error
	err := data.GetAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateCategory(t *testing.T) {
	// Create service and assign to data
	data := CategoryService{
		Items: []models.Category{{Id: "3", Name: "Nón dai"}},
	}

	// Execute method and if error happen send error
	err := data.Update()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
