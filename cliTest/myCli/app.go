package myCli

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

/**

postMan的form-data参数生成

*/

func MyCli() {

	app := &cli.App{
		Name:        "我的cli项目",
		Usage:       "常用工具集合",
		Description: "参数处理 ",
	}

	// 参数处理的cli
	params := Param()

	//

	app.Commands = append(app.Commands, params...)

	// 排序
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
