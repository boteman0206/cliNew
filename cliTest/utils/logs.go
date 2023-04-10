package utils

import (
	"path/filepath"
)

// 记录文件log信息
var (
	ERRLOG  *LogInfo
	INFOLOG *LogInfo
	DIR     string
	DIRLOG  string
	LOGTXT  = "log.txt"
)

func init() {

	DIR = GetDir()
	DIRLOG = filepath.Join(DIR, LOGTXT)

	CreateFile(DIRLOG)

	ERRLOG = NewErrLog()
	INFOLOG = NewINfoLog()
	INFOLOG.Message = "测试一下info: "
	ERRLOG.Message = "测试一下:error "

	INFOLOG.WriteLog()
	ERRLOG.WriteLog()

}

func NewErrLog() *LogInfo {

	return &LogInfo{Level: " ERROR", FilePath: DIRLOG}
}

func NewINfoLog() *LogInfo {

	return &LogInfo{Level: " INFO", FilePath: DIRLOG}
}

type LogInfo struct {
	FilePath string `json:"file_path"`
	Level    string `json:"level"`
	Message  string `json:"message"`
	Path     string `json:"path"`
}

func (l *LogInfo) WriteLog() {
	WriteFile(DIRLOG, l)
}
