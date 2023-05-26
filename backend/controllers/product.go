package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"commercial-shop.com/access"
)

func GetProductAll(c *gin.Context) {
	// Get limit
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 20
	}
	if limit <= 0 {
		limit = 20
	}

	// Get page
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	if page <= 0 {
		page = 1
	}

	data, err := access.FindProductAll(&limit, &page)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "can't get all product value"})
		return
	}
	c.JSON(http.StatusOK, data)
}
