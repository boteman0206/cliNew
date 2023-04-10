package test

// 子命令

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

/*
*
比较完整的一个cli程序示例
*/
func Test04() {

	app := &cli.App{
		Name:        "ZC-Web-Deploy",
		Usage:       "项目打包平台",
		Description: "项目打包平台,将打好的包上传至存储中心上面，运维人员作发布",
		Commands: []*cli.Command{
			{
				Name:    "project",
				Aliases: []string{"p"},
				Usage:   "deploy project",
				Subcommands: []*cli.Command{
					{
						Name:    "init",
						Aliases: []string{"i"},
						Usage:   "init the project",
						//./hello project b --domain test --tag iji
						Action: func(c *cli.Context) error {
							fmt.Println("我是initProject-----")
							return nil
						},
					},

					/*
						todo 最主要的是这个build命令，其中 flags是命令帮助的一个输出解释，比如说需要domain 和tag参数
						Action是命令的执行操作
					*/
					{
						Name:    "build",
						Aliases: []string{"b"},
						Usage:   "build the project as tar to ftp",
						Action: func(c *cli.Context) error {
							domain := c.String("domain")
							tag := c.String("tag")
							if domain == "" || tag == "" {
								fmt.Println("构建的项目域名或Tag不能为空")
								return nil
							}
							/*
								 ./hello project build --help  // 获取命令帮助
								./hello project b --domain test --tag iji  // 执行命令
							*/
							fmt.Println("我是打包项目build---domain is : ", domain, " tag is ", tag)
							return nil
						},
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "domain, d",
								Value:    "xxxx",
								Usage:    "指定工程的域名",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "tag, t",
								Value:    "release-tag",
								Usage:    "打包项目的git的tag发布版本号",
								Required: true,
							},
						},
					},
					{
						Name:    "list",
						Aliases: []string{"ls"},
						Usage:   "list the project",
						//./hello project list  执行命令

						Action: func(c *cli.Context) error {
							fmt.Println("暂无项目，可以初始化建立project init")
							return nil
						},
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("请指定具体操作: init、ls、build或用 -h 查看命令")
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
