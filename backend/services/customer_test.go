package services

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_CreateCustomer(t *testing.T) {
	// Create service and assign to data
	data := CustomerService{
		Items: []models.Customer{{
			Id:      "VIP03",
			Name:    "Khoa22",
			Phone:   "none",
			Address: "TP HCM",
		}},
	}

	// Execute method and if error happen send error
	err := data.Create()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteCustomer(t *testing.T) {
	// Create service and assign to data
	data := CustomerService{Items: []models.Customer{{
		Id: "VIP03",
	}}}

	// Execute method and if error happen send error
	err := data.Delete()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetCustomer(t *testing.T) {
	// Create service and assign to data
	data := CustomerService{Items: []models.Customer{{
		Id: "VIP03",
	}}}

	// Execute method and if error happen send error
	err := data.Get()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetAllCustomer(t *testing.T) {
	// Create limit, page, service and assign to data
	limit := 10
	page := 1
	data := CustomerService{}

	// Execute method and if error happen send error
	err := data.GetAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateCustomer(t *testing.T) {
	// Create service and assign to data
	data := CustomerService{
		Items: []models.Customer{{
			Id:      "VIP03",
			Name:    "Khoa22",
			Phone:   "0111432389",
			Address: "TP HCM",
		}},
	}

	// Execute method and if error happen send error
	err := data.Update()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
