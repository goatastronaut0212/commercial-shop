package routes

import (
	"github.com/gin-gonic/gin"

	"commercial-shop.com/controllers"
)

func addCategoryRoutes(rg *gin.RouterGroup) {
	route := rg.Group("category")

	route.GET("", controllers.GetAllCategory)
	route.GET("/:id", controllers.GetCategory)
	route.POST("", controllers.CreateCategory)
	route.PATCH("/:id", controllers.UpdateCategory)
	route.DELETE("/:id", controllers.DeleteCategory)
}
