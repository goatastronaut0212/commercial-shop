package category 

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
