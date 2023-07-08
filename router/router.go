package router

import (
	"github.com/gin-gonic/gin"
	"lendotopia.com/originator/middleware/auth"
	"lendotopia.com/originator/resumer"
)

func Init() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(auth.TokenAuthMiddleware())
	router.GET("/application/:id", resumer.LoadQuoteByApplicationID)
	router.Run("localhost:8080")
}
