package access

import (
	"testing"
	"time"

	"commercial-shop.com/models"
)

func Test_FindAllDiscount(t *testing.T) {
	limit := 10
	page := 10
	_, err := FindAllDiscount(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_FindDiscount(t *testing.T) {
	id := "XV1"
	_, err := FindDiscount(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_CreateDiscount(t *testing.T) {
	data := models.Discount{
		Id:          "TestN",
		Description: "Thông tin chưa cập nhập",
		Percent:     0.25,
		DateStart:   time.Now(),
		DateEnd:     time.Now(),
	}
	err := CreateDiscount(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateDiscount(t *testing.T) {
	data := models.Discount{
		Id:          "TestN",
		Description: "Test thử discount",
		Percent:     0.25,
		DateStart:   time.Now(),
		DateEnd:     time.Now(),
	}
	err := UpdateDiscount(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteDiscount(t *testing.T) {
	id := "TestN"
	err := DeleteDiscount(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
