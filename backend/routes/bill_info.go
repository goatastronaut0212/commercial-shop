package routes

import (
	"github.com/gin-gonic/gin"

	"commercial-shop.com/controllers"
)

func addBillInfoRoutes(rg *gin.RouterGroup) {
	route := rg.Group("billinfo")

	route.GET("", controllers.GetAllBill_Info)
	route.GET("/:id", controllers.GetBill_Info)
	route.POST("", controllers.CreateBillInfo)
	route.PATCH("/:id", controllers.UpdateBillInfo)
	route.DELETE("/:id", controllers.DeleteBillInfo)
}
