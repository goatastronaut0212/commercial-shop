package routes

import (
	"github.com/gin-gonic/gin"

	"commercial-shop.com/controllers"
)

func addBillDetailRoutes(rg *gin.RouterGroup) {
	route := rg.Group("bill-detail")

	route.GET("", controllers.GetAllBill_Detail)
	route.GET("/:id", controllers.GetBill_Detail)
	route.POST("", controllers.CreateBill_Detail)
	route.PATCH("/:id", controllers.UpdateBill_Detail)
	route.DELETE("/:id", controllers.DeleteBill_Detail)
}
