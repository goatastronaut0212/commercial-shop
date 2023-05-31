package access

import (
	"testing"

	"commercial-shop.com/models"
)

func Test_FindAllAccount(t *testing.T) {
	limit := 10
	page := 10
	_, err := FindAllAccount(&limit, &page)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_FindAccount(t *testing.T) {
	username := "003"
	_, err := FindAccount(&username)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_CreateAccount(t *testing.T) {
	data := models.Account{
		Username:    "Khoa999",
		CustomerId:  "VIP03",
		Password:    "1232",
		DisplayName: "Khoa",
		RoleId:      1,
	}
	err := CreateAccount(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_UpdateAccount(t *testing.T) {
	data := models.Account{
		Username:    "Khoa999",
		CustomerId:  "VIP03",
		Password:    "1232",
		DisplayName: "Khoa 1",
		RoleId:      1,
	}
	err := UpdateAccount(&data)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func Test_DeleteAccount(t *testing.T) {
	username := "Khoa999"
	err := DeleteAccount(&username)
	if err != nil {
		t.Fatalf("Err: %v", err)
	}
}
