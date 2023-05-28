package access

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_FindAllProductDetail(t *testing.T) {
	limit := 20
	page := 1
	_, err := FindAllProductDetail(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_FindProductDetail(t *testing.T) {
	id := "XD32"
	_, err := FindProductDetail(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_CreateProductDetail(t *testing.T) {
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

func Test_UpdateProductDetail(t *testing.T) {
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

func Test_DeleteProductDetail(t *testing.T) {
	id := "D39"
	err := DeleteProductDetail(&id)
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
}
