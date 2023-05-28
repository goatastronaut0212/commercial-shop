package access

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_FindAllProductImage(t *testing.T) {
	limit := 20
	page := 1
	_, err := FindAllProductImage(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_FindProductImage(t *testing.T) {
	id := "1"
	_, err := FindProductImage(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_CreateProductImage(t *testing.T) {
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

func Test_UpdateProductImage(t *testing.T) {
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

func Test_DeleteProductImage(t *testing.T) {
	id := "3"
	err := DeleteProductImage(&id)
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
}
