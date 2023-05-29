package routes

import (
	"github.com/gin-gonic/gin"

	"commercial-shop.com/controllers"
)

func addAccountRoutes(rg *gin.RouterGroup) {
	route := rg.Group("account")

	route.GET("", controllers.GetAllAccount)
	route.GET("/:id", controllers.GetAccount)
	route.POST("", controllers.CreateAccount)
	route.PATCH("/:id", controllers.UpdateAccount)
	route.DELETE("/:id", controllers.DeleteAccount)
}
