package chatlogger

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	kickchatwrapper "github.com/SongoMen/kick-chat-wrapper"
	"go.uber.org/zap"
)

const (
	channelsDataPath = "channels.json"
	logsPath         = "logs/"
)

func loadChannelsMetadata(zapLogger *zap.SugaredLogger) (map[int]string, error) {
	jsonFile, err := os.Open(channelsDataPath)
	if err != nil {
		zapLogger.Error(err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[int]string
	json.Unmarshal([]byte(byteValue), &result)

	return result, nil
}

func createLogDirectoryIfNotExists(fullPath string, zapLogger *zap.SugaredLogger) error {
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		err := os.MkdirAll(fullPath, 0750)
		if err != nil {
			zapLogger.Error(err)
			return err
		}
	}
	return nil
}

func saveLog(date time.Time, channelID int, channel string, user string, message string, badges string, zapLogger *zap.SugaredLogger) error {
	year, month, _ := date.Date()
	fullPath := filepath.Join(logsPath, strconv.Itoa(channelID), strconv.Itoa(year), strconv.Itoa(int(month)))

	if err := createLogDirectoryIfNotExists(fullPath, zapLogger); err != nil {
		return err
	}

	logFilePath := filepath.Join(fullPath, "logs.txt")
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0750)
	if err != nil {
		zapLogger.Error(err)
		return err
	}
	defer f.Close()

	logLine := date.String() + "\t" + strconv.Itoa(channelID) + "\t" + channel + "\t" + user + "\t" + message + "\t" + badges + "\n"
	if _, err := f.WriteString(logLine); err != nil {
		zapLogger.Error(err)
		return err
	}

	return nil
}

func StartLogger(zapLogger *zap.SugaredLogger) {
	channels, err := loadChannelsMetadata(zapLogger)
	if err != nil {
		zapLogger.Error(err)
		return
	}

	kickChatWrapper, err := kickchatwrapper.NewClient()
	if err != nil {
		zapLogger.Error(err)
		return
	}

	for id := range channels {
		kickChatWrapper.JoinChannelByID(id)
	}

	messages := kickChatWrapper.ListenForMessages()

	for message := range messages {
		if message.ChatroomID == 0 {
			continue
		}
		serializedBadges := ""
		for _, badge := range message.Sender.Identity.Badges {
			serializedBadges += badge.Text + ";"
		}
		if err := saveLog(message.CreatedAt, message.ChatroomID, channels[message.ChatroomID], message.Sender.Username, message.Content, serializedBadges, zapLogger); err != nil {
			zapLogger.Error(err)
		}
	}
}
