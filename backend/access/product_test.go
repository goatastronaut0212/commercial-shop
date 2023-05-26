package access

import (
	"testing"
	// _ "commercial-shop.com/models"
)

func TestProductFindAll(t *testing.T) {
	limit := 20
	page := 1
	_, err := FindProductAll(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
