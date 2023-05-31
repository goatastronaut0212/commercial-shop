package access

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_FindAllBillInfo(t *testing.T) {
	limit := 10
	page := 10
	_, err := FindAllBillInfo(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_FindBillInfo(t *testing.T) {
	id := "001"
	_, err := FindBillInfo(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_CreateBillInfo(t *testing.T) {
	data := models.BillInfo{
		Id:         "001",
		CustomerId: "VIP01",
		Date:       "22/04/2023",
		Status:     0,
		Payment:    1,
	}
	err := CreateBillInfo(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateBillInfo(t *testing.T) {
	data := models.BillInfo{
		Id:         "001",
		CustomerId: "VIP01",
		Date:       "22/04/2023",
		Status:     2,
		Payment:    1,
	}
	err := UpdateBillInfo(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteBillInfo(t *testing.T) {
	id := "001"
	err := DeleteBillInfo(&id)
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
}
