package myCli

import (
	"fmt"
	"my-cli/utils"
	"path/filepath"
	"testing"
)

func TestParam(t *testing.T) {

	dir := utils.DIR
	file := filepath.Join(dir, "a.txt")

	data := utils.ReadLine(file)

	fmt.Println(data)

}
