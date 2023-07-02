package main

import (
	"github.com/SongoMen/kickchatlogger/api"
	"github.com/SongoMen/kickchatlogger/chatlogger"
	"github.com/SongoMen/kickchatlogger/utils"
)

func main() {
	utils.InitializeLogger()
	utils.LoadChannelsMetadata()
	go chatlogger.StartLogger()
	api.StartServer()
}
