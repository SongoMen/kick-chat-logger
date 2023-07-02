package api

import (
	"github.com/SongoMen/kickchatlogger/utils"
	"github.com/gin-gonic/gin"
)

func GetUserLogs(c *gin.Context) {
	params := c.Request.URL.Query()
	utils.Logger.Info("GetUserLogs", params)
}
