package utils

import (
	"path/filepath"
	"strings"
)

// 记录文件log信息
var (
	DIR    string
	DIRLOG string
	LOGTXT = "log.txt"

	ERROR = "ERROR"
	INFO  = "INFO"
)

func init() {

	DIR = GetDir()
	DIRLOG = filepath.Join(DIR, LOGTXT)

	CreateFile(DIRLOG)

}

func NewLog() *LogInfo {
	return &LogInfo{FilePath: DIRLOG}
}

func (l *LogInfo) ERROR(str ...string) {
	l.Level = ERROR
	l.Message = ForMatMsg(str...)

	WriteFile(DIRLOG, l)
}

func (l *LogInfo) INFO(str ...string) {
	l.Level = INFO
	l.Message = ForMatMsg(str...)

	WriteFile(DIRLOG, l)
}

func ForMatMsg(str ...string) string {
	var strs = make([]string, len(str))
	for i := range str {
		strs = append(strs, str[i])
	}
	return strings.Join(strs, "")
}

type LogInfo struct {
	FilePath string `json:"file_path"`
	Level    string `json:"level"`
	Message  string `json:"message"`
	Path     string `json:"path"`
}
