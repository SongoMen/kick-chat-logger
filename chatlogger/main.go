package chatlogger

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	kickchatwrapper "github.com/SongoMen/kick-chat-wrapper"
	"github.com/SongoMen/kickchatlogger/utils"
)

const (
	logsPath = "logs/"
)

func createLogDirectoryIfNotExists(fullPath string) error {
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		err := os.MkdirAll(fullPath, 0750)
		if err != nil {
			utils.Logger.Error(err)
			return err
		}
	}
	return nil
}

func saveLog(date time.Time, channelID int, user string, message string, badges string) error {
	year, month, _ := date.Date()
	fullPath := filepath.Join(logsPath, "users", user, strconv.Itoa(channelID), strconv.Itoa(year), strconv.Itoa(int(month)))

	if err := createLogDirectoryIfNotExists(fullPath); err != nil {
		return err
	}

	logFilePath := filepath.Join(fullPath, "logs.txt")
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0750)
	if err != nil {
		utils.Logger.Error(err)
		return err
	}
	defer f.Close()

	logLine := date.String() + "\t" + user + "\t" + message + "\t" + badges + "\n"
	if _, err := f.WriteString(logLine); err != nil {
		utils.Logger.Error(err)
		return err
	}

	return nil
}

func StartLogger() {
	utils.Logger.Info("Logging Start for channels:", utils.Channels)

	kickChatWrapper, err := kickchatwrapper.NewClient()
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	for id := range utils.Channels {
		kickChatWrapper.JoinChannelByID(id)
	}

	messages := kickChatWrapper.ListenForMessages()

	for message := range messages {
		if message.ChatroomID == 0 || message.Content == "" {
			continue
		}
		serializedBadges := ""
		for _, badge := range message.Sender.Identity.Badges {
			serializedBadges += badge.Text + ";"
		}
		if err := saveLog(message.CreatedAt, message.ChatroomID, message.Sender.Username, message.Content, serializedBadges); err != nil {
			utils.Logger.Error(err)
		}
	}
}
