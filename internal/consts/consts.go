package consts

import (
	"os"
	"path/filepath"
	"strings"
)

const AppVersion = "1.0.3"
const AssetsVersion = 2

const DataDirectory = "data"
const DBFileName = "IP2LOCATION-LITE-DB3.BIN"

var Config struct {
	AccessLogFile string `description:"Access log file"`
	DataDir       string `description:"Application data directory"`
	DbUpdateTime  int64  `default:"60" description:"Delay in minutes between database reloading"`
	Deployment    string `default:"development" description:"Deployment type"`
	ErrorLogFile  string `description:"Error log file"`
	Host          string `default:"127.0.0.1" description:"Web server IP"`
	LimitRequests int    `default:"5" description:"Requests per second per one IP"`
	Port          string `default:"8080" description:"Web server port"`
	WebURL        string `default:"http://localhost:8080/" description:"Web server home URL"`
}

func DataPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return dir, err
	}
	return strings.Join(append([]string{dir}, DataDirectory), string(os.PathSeparator)), nil
}

func DataPathFile(filename ...string) (string, error) {
	dir, err := filepath.Abs(Config.DataDir)
	if err != nil {
		return dir, err
	}
	return strings.Join(append([]string{dir}, filename...), string(os.PathSeparator)), nil
}
