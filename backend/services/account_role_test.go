package services

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_GetAccountRole(t *testing.T) {
	// Create service and assign to data
	data := AccountRoleService{Items: []models.AccountRole{{
		Id: 3,
	}}}

	// Execute method and if error happen send error
	err := data.Get()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetAllAccountRole(t *testing.T) {
	// Create limit, page, service and assign to data
	limit := 10
	page := 1
	data := AccountRoleService{}

	// Execute method and if error happen send error
	err := data.GetAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_CreateAccountRole(t *testing.T) {
	// Create service and assign to data
	data := AccountRoleService{
		Items: []models.AccountRole{{
			Id:          3,
			Description: "Vô dụng",
		}},
	}

	// Execute method and if error happen send error
	err := data.Create()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateAccountRole(t *testing.T) {
	// Create service and assign to data
	data := AccountRoleService{
		Items: []models.AccountRole{{
			Id:          3,
			Description: "Vô dụng theo mặc định",
		}},
	}

	// Execute method and if error happen send error
	err := data.Update()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteAccountRole(t *testing.T) {
	// Create service and assign to data
	data := AccountRoleService{Items: []models.AccountRole{{
		Id: 3,
	}}}

	// Execute method and if error happen send error
	err := data.Delete()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
