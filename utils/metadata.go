package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	channelsDataPath = "channels.json"
)

type channelMetadata struct {
	Username string `json:"username"`
	StvID    string `json:"stvID"`
}

var ChannelsMetadata map[int]channelMetadata
var UserIdMapper map[string]int

func GetChannelID(channelName string) int {
	return UserIdMapper[channelName]
}

func LoadChannelsMetadata() {
	jsonFile, err := os.Open(channelsDataPath)
	if err != nil {
		Logger.Error(err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &ChannelsMetadata)
	UserIdMapper = make(map[string]int)
	for id, _ := range ChannelsMetadata {
		UserIdMapper[ChannelsMetadata[id].Username] = id
	}
}
