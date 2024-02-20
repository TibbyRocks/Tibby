package utils

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	Log        *slog.Logger
	LogLevel   *slog.LevelVar
	logDir     = os.Getenv("LOG_PATH")
	logOutputs io.Writer
)

func init() {
	//If the LOG_PATH environment variable is empty (non-existent) we just write to Stdout
	if logDir != "" {
		err := os.Mkdir(logDir, 0744)
		if err != nil && !os.IsExist(err) {
			fmt.Println("Failed to create log directory: ", err.Error())
		}
		LogFileName := fmt.Sprintf("%s/%s-%s.%s", logDir, "tibby", time.Now().Format("20060102"), "log.json")
		LogFile, err := os.OpenFile(LogFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Failed to open log file: ", err.Error())
		}
		logOutputs = io.MultiWriter(LogFile, os.Stdout)
	} else {
		logOutputs = os.Stdout
	}

	LogLevel = &slog.LevelVar{}
	opts := slog.HandlerOptions{
		Level: LogLevel,
	}
	handler := slog.NewJSONHandler(logOutputs, &opts)

	Log = slog.New(handler)
}

func LogCmd(i *discordgo.InteractionCreate) {
	user := GetUserobjectFromInteraction(i)
	Log.Info("Command used", "Command", i.ApplicationCommandData().Name, "Args", GetOptionsStringsFromInteraction(i), "Username", user.Username, "UserID", user.ID)
}
