package access

import (
	"testing"

	"commercial-shop.com/models"
	// _ "commercial-shop.com/models"
)

func TestProductFindAll(t *testing.T) {
	limit := 20
	page := 1
	_, err := FindProductAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestProductFindById(t *testing.T) {
	id := "1"
	_, err := FindCategoryById(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestProductCreate(t *testing.T) {
	data := models.Product{
		Id:         "3",
		IdCategory: "1",
		Name:       "Quần nam 3",
	}
	err := CreateProduct(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestProductUpdate(t *testing.T) {
	data := models.Product{
		Id:         "3",
		IdCategory: "1",
		Name:       "Quần nam 3",
	}
	err := UpdateProduct(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestProductDelete(t *testing.T) {
	id := "3"
	err := DeleteProduct(&id)
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
}
