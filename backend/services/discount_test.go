package services

import (
	"testing"
	"time"

	"commercial-shop.com/models"
)

func Test_CreateDiscount(t *testing.T) {
	// Create service and assign to data
	data := DiscountService{Items: []models.Discount{{
		Id:          "TestN",
		Description: "Thông tin chưa cập nhập",
		Percent:     0.25,
		DateStart:   time.Now(),
		DateEnd:     time.Now(),
	}}}

	// Execute method and if error happen send error
	err := data.Create()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteDiscount(t *testing.T) {
	// Create service and assign to data
	data := DiscountService{Items: []models.Discount{{
		Id: "TestN",
	}}}

	// Execute method and if error happen send error
	err := data.Delete()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetDiscount(t *testing.T) {
	// Create service and assign to data
	data := DiscountService{Items: []models.Discount{{
		Id: "XV1",
	}}}

	// Execute method and if error happen send error
	err := data.Get()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetAllDiscount(t *testing.T) {
	// Create limit, page, service and assign to data
	limit := 10
	page := 1
	data := DiscountService{}

	// Execute method and if error happen send error
	err := data.GetAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateDiscount(t *testing.T) {
	// Create service and assign to data
	data := DiscountService{Items: []models.Discount{{
		Id:          "TestN",
		Description: "Thông tin chưa cập nhập",
		Percent:     0.2,
		DateStart:   time.Now(),
		DateEnd:     time.Now(),
	}}}

	// Execute method and if error happen send error
	err := data.Update()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
