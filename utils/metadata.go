package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	channelsDataPath = "channels.json"
)

var Channels map[int]string

func GetChannelID(channelName string) int {
	channelID := 0
	for id, channel := range Channels {
		if channel == channelName {
			channelID = id
			break
		}
	}
	return channelID
}

func LoadChannelsMetadata() {
	jsonFile, err := os.Open(channelsDataPath)
	if err != nil {
		Logger.Error(err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &Channels)
}
