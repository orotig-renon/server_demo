package routers

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/orotig/server_demo/controllers"
)

func SetupRouter() *gin.Engine {

	gin.ForceConsoleColor()

	server := gin.New()
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("[%s][%s] \"%s %s %d %s\"\n",
			param.TimeStamp.Format(time.RFC1123),
			param.ClientIP,
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	}))
	server.Use(gin.Recovery())
	server.Use(cors.Default())

	v1 := server.Group("/api/v1")
	v1.Use()
	{
		quaranta := v1.Group("40")
		{
			quaranta.GET("/welder/order", controllers.ListWelderOrders)
			quaranta.POST("/welder/order", controllers.UpdateWeldersOrder)
			quaranta.GET("/marker/order", controllers.ListMarkerOrders)
			quaranta.GET("/marker/pending_order", controllers.ListPendingMarkerOrders)
			quaranta.POST("/marker/order", controllers.UpdateMarkerOrder)
		}
	}

	return server
}
