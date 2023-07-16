package api

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/SongoMen/kickchatlogger/utils"
	"github.com/gin-gonic/gin"
)

type UserLog struct {
	Date    string   `json:"date"`
	Message string   `json:"message"`
	Badges  []string `json:"badges"`
}

type LogResponse struct {
	Messages []UserLog `json:"messages"`
	Periods  []string  `json:"periods"`
}

func reverseSlice(s []UserLog) []UserLog {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func getMapKeys(m interface{}) []string {
	switch m := m.(type) {
	case map[string]string, map[string]int:
		v := reflect.ValueOf(m)
		keys := make([]string, 0, v.Len())
		for _, k := range v.MapKeys() {
			keys = append(keys, k.String())
		}
		return keys
	default:
		return nil
	}
}

func sortPeriods(periods []string) {
	sort.Slice(periods, func(i, j int) bool {
		a, _ := time.Parse("2006-1", periods[i])
		b, _ := time.Parse("2006-1", periods[j])
		return b.Before(a)
	})
}

func constructLogsPath(user string, channel int, period string) string {
	basePath := fmt.Sprintf("logs/users/%s/%d", user, channel)
	if period == "" {
		return basePath
	}
	date := strings.Split(period, "-")
	year, month := date[0], date[1]
	return filepath.Join(basePath, year, month, "logs.txt")
}

func processFileLogs(filePath string) []UserLog {
	result := make([]UserLog, 0)
	file, err := os.Open(filePath)
	if err != nil {
		utils.Logger.Error(err)
		return result
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splittedLine := strings.Split(scanner.Text(), "\t")
		result = append(result, UserLog{
			Date:    splittedLine[0],
			Message: splittedLine[1],
			Badges:  strings.Split(splittedLine[2], ";"),
		})
	}
	return reverseSlice(result)
}

func getUserPeriods(path string) (string, map[string]string, error) {
	latestPeriod := ""
	allPeriods := make(map[string]string)
	latestDate, err := time.Parse("2006-1", "0001-1")
	if err != nil {
		utils.Logger.Error("Error parsing date:", err)
		return "", nil, err
	}

	years, err := os.ReadDir(path)
	if err != nil {
		utils.Logger.Error("Error reading directory:", err)
		return "", nil, err
	}

	for _, year := range years {
		months, err := os.ReadDir(filepath.Join(path, year.Name()))
		if err != nil {
			utils.Logger.Error("Error reading directory:", err)
			continue
		}
		for _, month := range months {
			period := fmt.Sprintf("%s-%s", year.Name(), month.Name())

			date, err := time.Parse("2006-1", period)
			if err != nil {
				utils.Logger.Error("Error parsing date:", err)
				continue
			}

			if date.After(latestDate) {
				latestDate = date
				latestPeriod = period
			}

			allPeriods[period] = filepath.Join(path, year.Name(), month.Name(), "logs.txt")
		}
	}

	return latestPeriod, allPeriods, nil
}

func GetUserLogs(c *gin.Context) {
	params := c.Request.URL.Query()
	utils.Logger.Info("GetUserLogs", params)
	channelID := utils.GetChannelID(params.Get("channel"))
	if channelID == 0 || !params.Has("user") {
		c.JSON(200, LogResponse{Messages: []UserLog{}, Periods: []string{}})
		return
	}

	logsPath := constructLogsPath(params.Get("user"), channelID, "")
	latestPeriod, userPeriods, err := getUserPeriods(logsPath)
	if err != nil {
		c.JSON(200, LogResponse{Messages: []UserLog{}, Periods: []string{}})
		return
	}

	if params.Has("period") {
		if logPath, ok := userPeriods[params.Get("period")]; ok {
			periods := getMapKeys(userPeriods)
			sortPeriods(periods)
			response := LogResponse{
				Messages: processFileLogs(logPath),
				Periods:  periods,
			}
			c.JSON(200, response)
			return
		}
		c.JSON(400, gin.H{"error": "Period not found"})
		return
	}

	periods := getMapKeys(userPeriods)
	sortPeriods(periods)

	response := LogResponse{
		Messages: processFileLogs(userPeriods[latestPeriod]),
		Periods:  periods,
	}
	c.JSON(200, response)
}

func GetChannels(c *gin.Context) {
	c.JSON(200, getMapKeys(utils.UserIdMapper))
}

func getChannelMetadata(c *gin.Context) {
	params := c.Request.URL.Query()
	channelID := utils.GetChannelID(params.Get("channel"))
	if channelID == 0 {
		c.JSON(400, gin.H{"error": "Channel not found"})
		return
	}
	c.JSON(200, utils.ChannelsMetadata[channelID])
}
