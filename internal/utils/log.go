package utils

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"
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
		LogFileName := fmt.Sprintf("%s/%s-%s", logDir, "tibby", time.Now().Format("20060102"))
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
	handler := slog.NewTextHandler(logOutputs, &opts)

	Log = slog.New(handler)
}
