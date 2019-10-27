package main

import (
	"log"
	"os"
	"time"

	"github.com/tomoyamachi/imagecheck-for-gocon/pkg"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "nginxcheck"
	app.ArgsUsage = "image_name"

	app.Flags = []cli.Flag{
		cli.DurationFlag{
			Name:  "timeout, t",
			Value: time.Second * 90,
			Usage: "docker timeout. e.g) 5s, 5m...",
		},
		cli.StringFlag{
			Name:  "authurl",
			Usage: "registry authenticate url",
		},
		cli.StringFlag{
			Name:  "username",
			Usage: "registry login username",
		},
		cli.StringFlag{
			Name:  "password",
			Usage: "registry login password. Using --password via CLI is insecure.",
		},
		cli.BoolFlag{
			Name:  "insecure",
			Usage: "registry connect insecure",
		},
		cli.BoolTFlag{
			Name:  "nonssl",
			Usage: "registry connect without ssl",
		},
		cli.StringFlag{
			Name:  "cache-dir",
			Usage: "cache directory",
		},
	}
	app.Action = pkg.Run
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
