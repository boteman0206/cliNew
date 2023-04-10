package utils

import (
	"bufio"
	"cliNew/model"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/**
文件相关
*/

// 读取当前目录下的文件.txt
func ReadFile(fileName string) string {
	dir := GetDir()
	file := filepath.Join(dir, fileName)

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("read file fail", err)
		return ""
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}

	return string(fd)
}

func GetDir() string {

	dir, _ := os.Getwd()
	fmt.Println("当前路径：", dir)
	return dir

}

func CreateFile(name string) (*os.File, error) {
	// 创建一个新文件
	return os.Create(name)
}

func WriteFile(fileName string, log *LogInfo) {

	// 打开一个已存在的文件
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}

	// 在文件中写入数据
	str := fmt.Sprintf("Level: %s, Msg: %s \n", log.Level, log.Message)
	_, err = io.WriteString(file, string(str))
	if err != nil {
		fmt.Println("写入文件失败：", err)
		file.Close()
		return
	}

	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("关闭文件失败：", err)
		return
	}
}

func ReadLine(fileName string) []model.FormDataInfo {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)

	data := []model.FormDataInfo{}
	if err != nil {
		fmt.Println("Open file error!", err)
		return data
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			fmt.Println("读取到的line ： ", line)
			var values []string
			values = strings.Split(line, " ")
			if len(values) < 2 {
				values = strings.Split(line, "\n")
			}
			if len(values) < 2 {
				values = strings.Split(line, "\t")
			}
			if len(values) < 2 {
				values = strings.Split(line, "\r")
			}
			if len(values) < 2 {
				values = strings.Split(line, "\r\n")
			}

			var slicesStr []string
			for i := range values {
				v := values[i]
				if len(strings.TrimSpace(v)) > 0 {
					slicesStr = append(slicesStr, v)
				}
			}
			if len(slicesStr) == 2 {
				data = append(data, model.FormDataInfo{Key: slicesStr[0], Value: slicesStr[1]})
			}

		}
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return data
			}
		}
	}

	return data
}
