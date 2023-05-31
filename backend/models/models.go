package models

import "time"

type Account struct {
	Username    string `json:"username"      binding:"required"`
	CustomerId  string `json:"customerId"    binding:"required"`
	Password    string `json:"password"      binding:"required"`
	DisplayName string `json:"displayName"   binding:"required"`
	RoleId      int    `json:"roleID"        binding:"required"`
}

type BillInfo struct {
	Id         string `json:"id"            binding:"required"`
	CustomerId string `json:"customerId"    binding:"required"`
	Date       string `json:"billDate"      binding:"required"`
	Status     int    `json:"status"        binding:"required"`
	Payment    int    `json:"payment"       binding:"required"`
}

type BillDetail struct {
	Id         string `json:"id"             binding:"required"`
	BillId     string `json:"billId"         binding:"required"`
	ProductId  string `json:"productId"      binding:"required"`
	DiscountId string `json:"discountId"     binding:"required"`
	Amount     int    `json:"amount"         binding:"required"`
}

type Category struct {
	Id   string `json:"id"   binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Customer struct {
	Id      string `json:"id"         binding:"required"`
	Name    string `json:"name"       binding:"required"`
	Phone   string `json:"phone"      binding:"required"`
	Email   string `json:"email"      binding:"required"`
	Address string `json:"address"    binding:"required"`
}

type Discount struct {
	Id          string    `string:"id"          binding:"required"`
	Description string    `string:"description" binding:"required"`
	Percent     float64   `string:"percent"     binding:"required"`
	DateStart   time.Time `string:"dateStart"   binding:"required"`
	DateEnd     time.Time `string:"dateEnd"     binding:"required"`
}

type Product struct {
	Id         string `json:"id"         binding:"required"`
	IdCategory string `json:"idCategory" binding:"required"`
	Name       string `json:"name"       binding:"required"`
}

type ProductDetail struct {
	Id          string `json:"id"          binding:"required"`
	IdProduct   string `json:"idProduct"   binding:"required"`
	Color       string `json:"color"       binding:"required"`
	Fabric      string `json:"fabric"      binding:"required"`
	Size        string `json:"size"        binding:"required"`
	Price       int    `json:"price"       binding:"required"`
	Amount      int    `json:"amount"      binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ProductImage struct {
	Id              string `json:"id"              binding:"required"`
	IdProductDetail string `json:"idProductDetail" binding:"required"`
	Image           []byte `json:"image"           binding:"required"`
}
