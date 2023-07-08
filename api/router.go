package api

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(CORSMiddleware())

	v1api := router.Group("/api/v1")
	v1api.GET("/logs", GetUserLogs)
	v1api.GET("/channels", GetChannels)

	return router
}
