package services

import (
	"testing"
)

func Test_CreateCustomer(t *testing.T) {
	data := CustomerService{
		Items: []Customer{{
			Id:      "VIP03",
			Name:    "Khoa22",
			Phone:   "none",
			Email:   "jsjsjsjsjs@sasas.com",
			Address: "TP HCM",
		}},
	}
	err := data.Create()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateCustomer(t *testing.T) {
	data := CustomerService{
		Items: []Customer{{
			Id:      "VIP03",
			Name:    "Khoa22",
			Phone:   "0111432389",
			Email:   "jsjsjsjsjs@sasas.com",
			Address: "TP HCM",
		}},
	}
	err := data.Update()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetCustomer(t *testing.T) {
	data := CustomerService{Items: []Customer{{
		Id: "VIP03",
	}}}
	err := data.Get()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetAllCustomer(t *testing.T) {
	limit := 10
	page := 1
	data := CustomerService{}
	err := data.GetAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteCustomer(t *testing.T) {
	id := "VIP03"
	data := CustomerService{}
	err := data.Delete(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
