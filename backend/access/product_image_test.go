package access

import (
	"testing"

	"commercial-shop.com/models"
)

func TestProductImageFindAll(t *testing.T) {
	limit := 20
	page := 1
	_, err := FindProductImageAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestProductImageFindById(t *testing.T) {
	id := "1"
	_, err := FindProductImageById(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestProductImageCreate(t *testing.T) {
	data := models.ProductImage{
		Id:              "3",
		IdProductDetail: "XD32",
		Image:           nil,
	}
	err := CreateProductImage(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestProductImageUpdate(t *testing.T) {
	data := models.ProductImage{
		Id:              "3",
		IdProductDetail: "D33",
		Image:           nil,
	}
	err := UpdateProductImage(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestProductImageDelete(t *testing.T) {
	id := "3"
	err := DeleteProductImage(&id)
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
}
