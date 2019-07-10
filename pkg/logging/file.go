package logging

import (
	"fmt"
	"log"
	"os"
	"rehabilitation_prescription/pkg/setting"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

func mkDir() {
	dir, _ := os.Getwd()                                      // 返回当前目录的路径
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm) // 创建对应的目录以及子目录，对应mkdir -p; ModePerm = 0777
	if err != nil {
		panic(err)
	}
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("permission: %v", err)
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to open file, error: %v", err)
	}

	return file
}
