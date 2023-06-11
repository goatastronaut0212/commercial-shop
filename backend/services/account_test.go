package services

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_GetAccount(t *testing.T) {
	// Create service and assign to data
	data := AccountService{Items: []models.Account{{
		Username: "Khoa999",
	}}}

	// Execute method and if error happen send error
	err := data.Get()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetAllAccount(t *testing.T) {
	// Create limit, page, service and assign to data
	limit := 10
	page := 1
	data := AccountService{}

	// Execute method and if error happen send error
	err := data.GetAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_CreateAccount(t *testing.T) {
	// Create service and assign to data
	data := AccountService{
		Items: []models.Account{{
			Username:    "Khoa999",
			CustomerId:  "VIP03",
			Password:    "1232",
			DisplayName: "Khoa",
			RoleId:      1,
		}},
	}

	// Execute method and if error happen send error
	err := data.Create()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateAccount(t *testing.T) {
	// Create service and assign to data
	data := AccountService{
		Items: []models.Account{{
			Username:    "Khoa999",
			CustomerId:  "VIP03",
			Password:    "123",
			DisplayName: "Updated",
			RoleId:      1,
		}},
	}

	// Execute method and if error happen send error
	err := data.Update()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteAccount(t *testing.T) {
	// Create service and assign to data
	data := AccountService{Items: []models.Account{{
		Username: "Khoa999",
	}}}

	// Execute method and if error happen send error
	err := data.Delete()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
