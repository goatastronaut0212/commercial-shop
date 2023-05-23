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
    addCategoryRoutes(r) 
}
