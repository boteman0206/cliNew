package test

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

/**
支持标识位 flags
*/

func Test03() {

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "lang",                   // cli输入的参数  --lang 参数的值
			Value: "english",                // 参数的值  默认是englisg
			Usage: "通过输入的参数--lang的值蓝判断输出内容", // 命令解释说明
		},
	}

	app.Action = func(ctx *cli.Context) error {

		name := "huanhuan"
		if ctx.NArg() > 0 {
			name = ctx.Args().Get(0)
		}

		if ctx.String("lang") == "english" {
			fmt.Println("english -- Hala", name)
		} else {
			fmt.Println("other -- Hello", name)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

/**
./hello   输出english -- Hala huanhuan

加上参数
./hello -lang chanese  输出 other -- Hello huanhuan

*/
