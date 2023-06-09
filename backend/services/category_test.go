package services

import (
	"testing"
)

func Test_CreateCategory(t *testing.T) {
	data := CategoryService{
		Items: []Category{{Id: "3", Name: "Nón"}},
	}
	err := data.Create()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateCategory(t *testing.T) {
	data := CategoryService{
		Items: []Category{{Id: "3", Name: "Nón dai"}},
	}
	err := data.Update()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetCategory(t *testing.T) {
	data := CategoryService{Items: []Category{{
		Id: "1",
	}}}
	err := data.Get()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_GetAllCategory(t *testing.T) {
	limit := 10
	page := 1
	data := CategoryService{}
	err := data.GetAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteCategory(t *testing.T) {
	data := CategoryService{Items: []Category{{
		Id: "3",
	}}}
	err := data.Delete()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
