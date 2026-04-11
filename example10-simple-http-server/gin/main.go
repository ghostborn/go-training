package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ginEngine() *gin.Engine {
	router := gin.Default()

	router.GET("/welcome", func(ctx *gin.Context) {
		firstname := ctx.DefaultQuery("firstname", "Guest")
		lastname := ctx.Query("lastname")
		ctx.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": gin.H{"data1": "data2"}})
	})

	return router
}

func main() {
	r := ginEngine()
	r.Run(":8080")
}
