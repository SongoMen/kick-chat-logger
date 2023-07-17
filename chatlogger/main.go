package chatlogger

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	kickchatwrapper "github.com/SongoMen/kick-chat-wrapper"
	"github.com/SongoMen/kickchatlogger/utils"
)

const (
	logsPath = "logs/"
)

func cleaner(text string) string {
	const replacement = ""
	var replacer = strings.NewReplacer(
		"\r\n", replacement,
		"\r", replacement,
		"\n", replacement,
		"\v", replacement,
		"\f", replacement,
		"\u0085", replacement,
		"\u2028", replacement,
		"\u2029", replacement,
	)
	return replacer.Replace(text)
}

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

func saveLog(date time.Time, channelID int, sender kickchatwrapper.Sender, message string, badges string) error {
	const TimeFormat string = "%04d/%02d"
	periodPath := fmt.Sprintf(TimeFormat, date.Year(), date.Month())
	fullPath := filepath.Join(logsPath, "users", strings.ToLower(sender.Username), strconv.Itoa(channelID), periodPath)

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

	logLine := date.String() + "\t" + cleaner(message) + "\t" + badges + "\t" + sender.Identity.Color + "\n"
	if _, err := f.WriteString(logLine); err != nil {
		utils.Logger.Error(err)
		return err
	}

	return nil
}

func StartLogger() {
	utils.Logger.Info("Logging Start for channels:", utils.ChannelsMetadata)

	kickChatWrapper, err := kickchatwrapper.NewClient()
	kickChatWrapper.SetDebug(true)
	if err != nil {
		utils.Logger.Error(err)
		return
	}

	for id := range utils.ChannelsMetadata {
		kickChatWrapper.JoinChannelByID(id)
	}

	messages := kickChatWrapper.ListenForMessages()

	for message := range messages {
		if message.ChatroomID == 0 || message.Content == "" {
			continue
		}
		serializedBadges := ""
		badgesLength := len(message.Sender.Identity.Badges)
		for index, badge := range message.Sender.Identity.Badges {
			serializedBadges += badge.Text
			if index != badgesLength-1 {
				serializedBadges += ";"
			}
		}
		if err := saveLog(message.CreatedAt, message.ChatroomID, message.Sender, message.Content, serializedBadges); err != nil {
			utils.Logger.Error(err)
		}
	}
}
