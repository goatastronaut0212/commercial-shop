package models

type Category struct {
	Id   string `json:"id"   binding:"required"`
	Name string `json:"name" binding:"required"`
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
