package test

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func Test01() {
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
