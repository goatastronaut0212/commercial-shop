package routes 

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"

    "commercial-shop.com/access/category"
)

func addCategoryRoutes(rg *gin.RouterGroup) {
    categoryRoute := rg.Group("/category")

    categoryRoute.GET("/", func (c *gin.Context) {
        data, err := category.FindAll()
        
        if err != nil {
            log.Fatal(err)
        }
        
        c.JSON(http.StatusOK, data)
    })
    categoryRoute.GET("/:id", func (c *gin.Context) {
        data, err := category.FindById(c.Params.ByName("id"))
        
        if err != nil {

        }

        c.JSON(http.StatusOK, data)
    })
}
