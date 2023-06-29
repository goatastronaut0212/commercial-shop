package services

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_CreateProductDetail(t *testing.T) {
	// Create service and assign to data
	data := ProductDetailService{Items: []models.ProductDetail{{
		Id:          "D39",
		IdProduct:   "1",
		Color:       "Đen",
		Fabric:      "Vải",
		Size:        "39",
		Price:       165000,
		Amount:      100,
		Description: "Thông tin chưa cập nhập",
	}}}

	// Execute method and if error happen send error
	err := data.Create()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteProductDetail(t *testing.T) {
	// Create service and assign to data
	data := ProductDetailService{Items: []models.ProductDetail{{
		Id: "D39",
	}}}

	// Execute method and if error happen send error
	err := data.Delete()
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
}

func Test_GetProductDetail(t *testing.T) {
	// Create service and assign to data
	data := ProductDetailService{Items: []models.ProductDetail{{
		Id: "XD32",
	}}}

	// Execute method and if error happen send error
	err := data.Get()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetAllProductDetail(t *testing.T) {
	// Create limit, page, service and assign to data
	limit := 20
	page := 1
	data := ProductDetailService{}

	// Execute method and if error happen send error
	err := data.GetAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateProductDetail(t *testing.T) {
	// Create service and assign to data
	data := ProductDetailService{Items: []models.ProductDetail{{
		Id:          "D39",
		IdProduct:   "1",
		Color:       "Đen",
		Fabric:      "Vải",
		Size:        "39",
		Price:       165000,
		Amount:      90,
		Description: "Thông tin chưa cập nhập",
	}}}

	// Execute method and if error happen send error
	err := data.Update()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
