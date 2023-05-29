package models

import "time"

type Category struct {
	Id   string `json:"id"   binding:"required"`
	Name string `json:"name" binding:"required"`
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


type Customer struct {
	Id      string `json:"id"         binding:"required"`
	Name    string `json:"name"       binding:"required"`
	Phone   string `json:"phone"      binding:"required"`
	Email   string `json:"email"      binding:"required"`
	Address string `json:"address"    binding:"required"`
}

type Account struct {
	UserName   string `json:"username"      binding:"required"`
	CustomerID string `json:"CustomerID"    binding:"required"`
	PassWord   string `json:"password"      binding:"required"`
	DiplayName string `json:"displayname"   binding:"required"`
	RoleID     int    `json:"roleID"        binding:"required"`
}

type Bill_Info struct{
	ID         string `json:"id"            binding:"required"`
	CustomerID string `json:"CustomerID"    binding:"required"`
	BillDate   string `json:"BillDate"      binding:"required"`
}

type Bill_Detail struct{
	ID          string `json:"id"             binding:"required"`
	BillID      string `json:"BillID"         binding:"required"`
	ProductID   string `json:"ProductID"      binding:"required"`
	DiscountID	string `json:"DiscountID"     binding:"required"`
}