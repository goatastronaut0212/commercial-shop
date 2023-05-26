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
	i := "1"
	_, err := FindCategoryById(&i)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestCategoryCreate(t *testing.T) {
	c := models.Category{Id: "3", Name: "Nón"}
	err := CreateCategory(&c)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestCategoryUpdate(t *testing.T) {
	c := models.Category{Id: "3", Name: "Áo khoác"}
	err := UpdateCategory(&c)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestCategoryDelete(t *testing.T) {
	i := "3"
	err := DeleteCategory(&i)
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
}
