package category_test

import (
	"testing"

	"commercial-shop.com/access"
	"commercial-shop.com/models"
)

func TestCategoryFindAll(t *testing.T) {
	limit := 10
	page := 10
	_, err := access.FindCategoryAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestCategoryFindById(t *testing.T) {
	i := "1"
	_, err := access.FindCategoryById(&i)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestCategoryCreate(t *testing.T) {
	c := models.Category{Id: "3", Name: "Nón"}
	err := access.CreateCategory(&c)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestCategoryUpdate(t *testing.T) {
	c := models.Category{Id: "3", Name: "Áo khoác"}
	err := access.UpdateCategory(&c)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestCategoryDelete(t *testing.T) {
	i := "3"
	err := access.DeleteCategory(&i)
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
}
