package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"main/modules/mysql"
	s "main/modules/structs"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func Api() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/api/table/:name", func(c *gin.Context) {
		//get body
		body := c.Request.Body
		value, _ := ioutil.ReadAll(body)
		var data map[string]string
		json.Unmarshal([]byte(value), &data)

		//insert data
		var requestData s.RequestTableData
		requestData.StartDate = data["startDate"]
		requestData.EndDate = data["endDate"]
		requestData.Name = c.Param("name")

		//response
		if requestData.Name == "seller" {
			c.JSON(http.StatusOK, gin.H{
				"data": mysql.GetSellerTable(requestData),
			})
		} else if requestData.Name == "order" {
			c.JSON(http.StatusOK, gin.H{
				"data": mysql.GetOrderTable(requestData),
			})
		}
	})

	router.POST("/api/chart/:name", func(c *gin.Context) {
		//get body
		body := c.Request.Body
		value, _ := ioutil.ReadAll(body)
		var data map[string]string
		json.Unmarshal([]byte(value), &data)

		//insert data
		var requestData s.RequestChartData
		requestData.Cycle = data["cycle"]
		requestData.Event = data["event"]
		requestData.Segment = data["segment"]
		requestData.StartDate = data["startDate"]
		requestData.EndDate = data["endDate"]
		requestData.Name = c.Param("name")
		//response
		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetChart(requestData),
		})
	})

	router.GET("/api/products/:limit", func(c *gin.Context) {
		var limit string = c.Param("limit")
		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetProducts(limit),
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
