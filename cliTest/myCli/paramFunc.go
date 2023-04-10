package myCli

import (
	"cliNew/utils"
	"fmt"
	"path/filepath"
	"time"
)

//param相关的方法

/*
*
处理文件参数内容
*/
func disposeContent(fileName string) {

	dir := utils.DIR
	file := filepath.Join(dir, fileName)

	data := utils.ReadLine(file)

	fmt.Println("读取的data： ", data)

	fileForm := filepath.Join(dir, fmt.Sprintf("form_data_%d.txt", time.Now().UnixNano()))
	fd, err := utils.CreateFile(fileForm)
	if err != nil {
		defer fd.Close()
	}
	//写入文件当中
	for i := range data {

		newStr := data[i].Key + ":" + data[i].Value + "\r\n"
		fd.WriteString(newStr)

	}

}
