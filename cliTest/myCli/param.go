package myCli

/**
参数相关处理的cli程序

*/
import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
)

/*
*
比较完整的一个cli程序示例
*/
func Param() []*cli.Command {

	return []*cli.Command{

		//./cliTest param   f -f a.txt
		{
			Name:    "param",
			Aliases: []string{"p"},
			Usage:   "param handling",
			Subcommands: []*cli.Command{

				// formData 处理formData的参数，使用抓包工具的时候拿不到formData的数据，需要自己一个个组装
				{
					Name:    "formData",
					Aliases: []string{"f"},
					Usage:   "读取当前目录下的txt文件，将抓包工具的form-data参数格式化处理，然后可以直接使用",
					Action: func(c *cli.Context) error {
						fileName := c.String("fileName")
						if len(strings.TrimSpace(fileName)) <= 0 {
							fmt.Println("请指定文件名称.txt文件")
							return nil
						}
						if !strings.HasSuffix(fileName, ".txt") {
							fmt.Println("文件名称需为.txt文件")
							return nil
						}
						fmt.Println("参数 is : ", fileName)
						disposeContent(fileName)
						return nil
					},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "fileName",
							Aliases:  []string{"f"},
							Value:    "formData.txt",
							Usage:    "指定当前文件的名称",
							Required: true,
						},
					},
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("请指定具体操作: formData 或用 -h 查看命令")
				return nil
			},
		},
	}

}
