package access

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_FindAllCategory(t *testing.T) {
	limit := 10
	page := 10
	_, err := FindAllCategory(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_FindCategory(t *testing.T) {
	id := "1"
	_, err := FindCategory(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_CreateCategory(t *testing.T) {
	data := models.Category{Id: "3", Name: "Nón"}
	err := CreateCategory(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateCategory(t *testing.T) {
	data := models.Category{Id: "3", Name: "Áo khoác"}
	err := UpdateCategory(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteCategory(t *testing.T) {
	id := "3"
	err := DeleteCategory(&id)
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
}
