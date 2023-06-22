package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run(port string) {
	getRoutes()
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE"}
	router.Use(cors.New(config))
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
