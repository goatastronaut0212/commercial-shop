package access

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_FindAllBillDetail(t *testing.T) {
	limit := 10
	page := 10
	_, err := FindAllBillDetail(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_FindBillDetail(t *testing.T) {
	id := "003"
	_, err := FindBillDetail(&id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_CreateBillDetail(t *testing.T) {
	data := models.BillDetail{
		Id:         "003",
		BillId:     "001",
		ProductId:  "3",
		DiscountId: "1",
		Amount:     2,
	}
	err := CreateBillDetail(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateBillDetail(t *testing.T) {
	data := models.BillDetail{
		Id:         "003",
		BillId:     "001",
		ProductId:  "3",
		DiscountId: "1",
		Amount:     3,
	}
	err := UpdateBillDetail(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteBillDetail(t *testing.T) {
	id := "003"
	err := DeleteBillDetail(&id)
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
}
