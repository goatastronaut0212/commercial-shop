package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"

    "commercial-shop.com/access/category" 
)

func GetCategoryAll(c *gin.Context) {
    data, err := category.FindAll()
     
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"data": "can't get all category values"})
    }
     
    err = c.ShouldBindJSON(data)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }

    c.JSON(http.StatusOK, data)
}

func GetCategoryById(c *gin.Context) {
    data, err := category.FindById(c.Params.ByName("id"))
        
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"data": "can't get category value"})
    }

    c.JSON(http.StatusOK, data)
}
