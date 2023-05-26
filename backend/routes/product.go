package routes

import (
	"github.com/gin-gonic/gin"

	"commercial-shop.com/controllers"
)

func addProductRoutes(rg *gin.RouterGroup) {
	categoryRoute := rg.Group("product")

	categoryRoute.GET("", controllers.GetProductAll)
}
