package myCli

import (
	"cliNew/utils"
	"fmt"
	"path/filepath"
	"testing"
)

func TestParam(t *testing.T) {

	dir := utils.DIR
	file := filepath.Join(dir, "a.txt")

	data := utils.ReadLine(file)

	fmt.Println(data)

}
