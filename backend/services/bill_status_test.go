package services

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_GetBillStatus(t *testing.T) {
	// Create service and assign to data
	data := BillStatusService{Items: []models.BillStatus{{
		Id: 3,
	}}}

	// Execute method and if error happen send error
	err := data.Get()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetAllBillStatus(t *testing.T) {
	// Create limit, page, service and assign to data
	limit := 10
	page := 1
	data := BillStatusService{}

	// Execute method and if error happen send error
	err := data.GetAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_CreateBillStatus(t *testing.T) {
	// Create service and assign to data
	data := BillStatusService{
		Items: []models.BillStatus{{
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

func Test_UpdateBillStatus(t *testing.T) {
	// Create service and assign to data
	data := BillStatusService{
		Items: []models.BillStatus{{
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

func Test_DeleteBillStatus(t *testing.T) {
	// Create service and assign to data
	data := BillStatusService{Items: []models.BillStatus{{
		Id: 3,
	}}}

	// Execute method and if error happen send error
	err := data.Delete()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
