package routes 

import (
    "github.com/gin-gonic/gin"

    "commercial-shop.com/controllers"
)

func addCategoryRoutes(rg *gin.RouterGroup) {
    categoryRoute := rg.Group("category")

    categoryRoute.GET("", controllers.GetCategoryAll)
    categoryRoute.GET("/:id", controllers.GetCategoryById)
    categoryRoute.POST("", controllers.CreateCategory)
    categoryRoute.PATCH("/:id", controllers.UpdateCategory)
    categoryRoute.DELETE("/:id", controllers.DeleteCategory)
}
