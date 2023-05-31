package access

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_FindAllCustomer(t *testing.T) {
	limit := 10
	page := 10
	_, err := FindAllCustomer(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_FindCustomer(t *testing.T) {
	id := "XV1"
	_, err := FindCustomer(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_CreateCustomer(t *testing.T) {
	data := models.Customer{
		Id:      "VIP03",
		Name:    "Khoa22",
		Phone:   "none",
		Email:   "jsjsjsjsjs@sasas.com",
		Address: "TP HCM",
	}
	err := CreateCustomer(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateCustomer(t *testing.T) {
	data := models.Customer{
		Id:      "VIP03",
		Name:    "Khoa22",
		Phone:   "12345",
		Email:   "jsjsjsjsjs@sasas.com",
		Address: "TP HCM",
	}
	err := UpdateCustomer(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteCustomer(t *testing.T) {
	id := "VIP03"
	err := DeleteCustomer(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
