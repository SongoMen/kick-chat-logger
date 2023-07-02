package api

import "github.com/SongoMen/kickchatlogger/utils"

func StartServer() {
	router := InitRouter()
	utils.Logger.Info("Starting server on port 8080")
	router.Run(":8080")
}
