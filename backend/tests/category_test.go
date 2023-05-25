package category_test 

import (
    "testing"
)

func TestFindAll(t *testing.T) {
    _, err := FindAll()

    if err != nil {
        t.Fatalf("Error: %v", err)
    }
}

func TestFindById(t *testing.T) {
    _, err := FindById("1")

    if err != nil {
        t.Fatalf("Error: %v", err)
    }
}

func TestCreate(t *testing.T) {
    err := Create("3", "Nón")
    if err != nil {
        t.Fatalf("Error: %v", err)
    }
}

func TestUpdate(t *testing.T) {
    err := Update("3", "Áo khoác")
    if err != nil {
        t.Fatalf("Error: %v", err)
    }
}

func TestDelete(t *testing.T) {
    err := Delete("3")
    if err != nil {
        t.Fatalf("Err: %v", err)
    }
}
