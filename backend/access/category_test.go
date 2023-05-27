package access

import (
	"testing"

	"commercial-shop.com/models"
)

func TestCategoryFindAll(t *testing.T) {
	limit := 10
	page := 10
	_, err := FindCategoryAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestCategoryFindById(t *testing.T) {
	id := "1"
	_, err := FindCategoryById(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestCategoryCreate(t *testing.T) {
	data := models.Category{Id: "3", Name: "Nón"}
	err := CreateCategory(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestCategoryUpdate(t *testing.T) {
	data := models.Category{Id: "3", Name: "Áo khoác"}
	err := UpdateCategory(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestCategoryDelete(t *testing.T) {
	id := "3"
	err := DeleteCategory(&id)
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
}
