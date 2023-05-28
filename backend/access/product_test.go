package access

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_FindAllProduct(t *testing.T) {
	limit := 20
	page := 1
	_, err := FindAllProduct(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_FindProduct(t *testing.T) {
	id := "1"
	_, err := FindProduct(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_CreateProduct(t *testing.T) {
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

func Test_UpdateProduct(t *testing.T) {
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

func Test_DeleteProduct(t *testing.T) {
	id := "3"
	err := DeleteProduct(&id)
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
}
