package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/orotig/server_demo/models"
)

func ListWelderOrders(c *gin.Context) {
	var welder_orders []models.WelderTest

	// var response models.MarkerTestResponse

	// response.OrotigTask = marker_order
	err := models.GetWelderOrders(&welder_orders)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, welder_orders)
	}
}

func UpdateWeldersOrder(c *gin.Context) {
	var welder_order models.WelderTest
	c.BindJSON(&welder_order)

	fmt.Println(welder_order.StartDate)

	err := models.UpdateWelderOrder(&welder_order)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, welder_order)
	}
}
