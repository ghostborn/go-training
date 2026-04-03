package main

import (
	"github.com/gin-gonic/gin"
)

func ginEngine() *gin.Engine{
	router := gin.Default()
	
	router.GET("")
}