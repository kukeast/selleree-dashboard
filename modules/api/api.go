package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"main/modules/mysql"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func Api() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/api/products/:limit", func(c *gin.Context) {
		var limit string = c.Param("limit")
		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetProducts(limit),
		})
	})

	router.POST("/api/orders/:limit", func(c *gin.Context) {
		var limit string = c.Param("limit")

		body := c.Request.Body
		value, _ := ioutil.ReadAll(body)
		var data map[string]string
		json.Unmarshal([]byte(value), &data)

		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetOrders(limit, data["sortBy"]),
		})
	})

	router.GET("/api/shopggus", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetShopggus(),
		})
	})

	router.GET("/api/today", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetToday(),
		})
	})

	router.GET("/api/today-chart/:name", func(c *gin.Context) {
		var name string = c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetTodayChart(name),
		})
	})

	router.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
