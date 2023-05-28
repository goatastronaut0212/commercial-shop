package routes

import (
	"github.com/gin-gonic/gin"

	"commercial-shop.com/controllers"
)

func addCustomerRoutes(rg *gin.RouterGroup) {
	route := rg.Group("customer")

	route.GET("", controllers.GetAllCustomer)
	route.GET("/:id", controllers.GetCustomer)
	route.POST("", controllers.CreateCustomer)
	// route.PATCH("/:id", controllers.UpdateProduct)
	// route.DELETE("/:id", controllers.DeleteProduct)
}
