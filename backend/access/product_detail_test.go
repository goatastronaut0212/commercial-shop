package access

import (
	"testing"

	"commercial-shop.com/models"
)

func TestProductDetailFindAll(t *testing.T) {
	limit := 20
	page := 1
	_, err := FindProductDetailAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestProductDetailFindById(t *testing.T) {
	id := "XD32"
	_, err := FindProductDetailById(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestProductDetailCreate(t *testing.T) {
	data := models.ProductDetail{
		Id:          "D39",
		IdProduct:   "1",
		Color:       "Đen",
		Fabric:      "Vải",
		Size:        "39",
		Price:       165000,
		Amount:      100,
		Description: "Thông tin chưa cập nhập",
	}
	err := CreateProductDetail(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestProductDetailUpdate(t *testing.T) {
	data := models.ProductDetail{
		Id:          "D39",
		IdProduct:   "1",
		Color:       "Đen",
		Fabric:      "Vải",
		Size:        "39",
		Price:       165000,
		Amount:      90,
		Description: "Thông tin chưa cập nhập",
	}
	err := UpdateProductDetail(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestProductDetailDelete(t *testing.T) {
	id := "D39"
	err := DeleteProductDetail(&id)
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
}
