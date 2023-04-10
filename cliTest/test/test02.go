package test

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Test02() {
	app := &cli.App{
		Name:  "hello",
		Usage: "hello world example",
		Action: func(c *cli.Context) error {
			fmt.Println("-----hello world-----")
			n := c.NArg()
			for i := 0; i < n; i++ {
				fmt.Println("获取的参数： ", c.Args().Get(i))
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
