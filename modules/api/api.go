package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"main/modules/mysql"
	"main/modules/oauth"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func Api() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/api/cover", TokenAuthMiddleware(), func(c *gin.Context) {
		var id string = c.Query("id")
		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetCover(id),
		})
	})

	router.GET("/api/products", TokenAuthMiddleware(), func(c *gin.Context) {
		var limit string = c.Query("limit")
		var id string = c.Query("id")
		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetProducts(limit, id),
		})
	})

	router.GET("/api/orders", TokenAuthMiddleware(), func(c *gin.Context) {
		var limit string = c.Query("limit")
		var sortBy string = c.Query("sortBy")
		var id string = c.Query("id")
		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetOrders(limit, sortBy, id),
		})
	})

	router.GET("/api/order/:orderId", TokenAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetOrderDetail(c.Param("orderId")),
		})
	})

	router.GET("/api/shopggus", TokenAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetShopggus(),
		})
	})

	router.GET("/api/today", TokenAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetToday(),
		})
	})

	router.GET("/api/today-chart/:name", TokenAuthMiddleware(), func(c *gin.Context) {
		var name string = c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetTodayChart(name),
		})
	})

	router.POST("/api/funnel", TokenAuthMiddleware(), func(c *gin.Context) {
		body := c.Request.Body
		value, _ := ioutil.ReadAll(body)
		var data map[string]string
		json.Unmarshal([]byte(value), &data)

		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetFunnel(data["startDate"], data["endDate"]),
		})
	})

	router.POST("/api/payment-setting", TokenAuthMiddleware(), func(c *gin.Context) {
		body := c.Request.Body
		value, _ := ioutil.ReadAll(body)
		var data map[string]string
		json.Unmarshal([]byte(value), &data)

		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetPaymentSetting(data["startDate"], data["endDate"]),
		})
	})

	router.POST("/api/sellers/:limit", TokenAuthMiddleware(), func(c *gin.Context) {
		var limit string = c.Param("limit")
		body := c.Request.Body
		value, _ := ioutil.ReadAll(body)
		var data map[string]string
		json.Unmarshal([]byte(value), &data)

		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetSellers(data["startDate"], data["endDate"], data["segment"], limit),
		})
	})
	router.GET("/api/seller/:id", TokenAuthMiddleware(), func(c *gin.Context) {
		var id string = c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"data": mysql.GetSeller(id),
		})
	})

	router.POST("/api/login", oauth.LogIn)
	router.GET("/api/test", TokenAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "test gogo",
		})
	})
	router.POST("api/refresh", oauth.Refresh)
	router.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST")
		c.Header("Access-Control-Expose-Headers", "Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := oauth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
