package routes

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run(port string) {
	getRoutes()
	router.Run(":" + port)
}

func getRoutes() {
	r := router.Group("/")
	addAccountRoutes(r)
	addAccountRoleRoutes(r)
	addBillDetailRoutes(r)
	addBillInfoRoutes(r)
	addBillStatusRoutes(r)
	addCategoryRoutes(r)
	addCustomerRoutes(r)
	addDiscountRoutes(r)
	addProductRoutes(r)
	addProductDetailRoutes(r)
	addProductImageRoutes(r)
}
