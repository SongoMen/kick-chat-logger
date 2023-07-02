package api

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	v1api := router.Group("/api/v1")
	v1api.GET("/logs", GetUserLogs)

	return router
}
