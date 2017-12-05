package main

import (
	"os"

	"github.com/wzshiming/pot/cli"
	"github.com/wzshiming/pot/cmd/pot/generate"
	"github.com/wzshiming/pot/cmd/pot/logo"
	"github.com/wzshiming/pot/cmd/pot/run"
)

func main() {

	app := cli.App{}

	app.Name = "pot"

	app.Usage = "Pot is web framework for the Go programming language."

	app.Version = "0.1"

	app.Authors = []*cli.Author{
		{
			Name:  "wzshiming",
			Email: "wzshiming@foxmail.com",
		},
	}

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "hidden-logo",
			Aliases: []string{"hl"},
			Usage:   "hidden logo",
			Value:   false,
		},
	}

	app.Before = func(c *cli.Context) error {
		if !c.Bool("hidden-logo") {
			logo.PrintLogo("V0.1")
		}
		return nil
	}

	app.Commands = Commands()

	app.Run(os.Args)
}

func Commands() []*cli.Command {
	r := []*cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Run pot service",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "gg",
					Usage: "go generate ./...",
					Value: false,
				},
			},
			Action: func(c *cli.Context) error {
				b := c.Bool("gg")
				ff := func() {
					run.GoBuild(func() {
						if b {
							run.GoGenerate()
						}
					})
				}
				ff()
				return run.NewWatcher([]string{"./"}, ff)
			},
		},
		{
			Name:        "generate",
			Aliases:     []string{"gen", "g"},
			Usage:       "Generate commands",
			Subcommands: SubcommandsGenerate(),
		},
	}
	return r
}

func SubcommandsGenerate() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "doc",
			Aliases: []string{"d"},
			Usage:   "Generate doc",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "controllers",
					Aliases: []string{"c"},
					Usage:   "Controllers dir",
					Value:   "./controllers",
				},
				&cli.StringFlag{
					Name:    "routers",
					Aliases: []string{"r"},
					Usage:   "Routers dir",
					Value:   "./routers",
				},
				&cli.StringFlag{
					Name:    "swagger",
					Aliases: []string{"s"},
					Usage:   "Swagger config",
					Value:   "./swagger",
				},
			},
			Action: func(c *cli.Context) error {
				generate.GenerateDocs(c.String("routers"), c.String("controllers"), c.String("swagger"))
				return nil
			},
		},
		{
			Name:    "router",
			Aliases: []string{"rou", "r"},
			Usage:   "Generate router",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "controllers",
					Aliases: []string{"c"},
					Usage:   "Controllers dir",
				},
				&cli.StringFlag{
					Name:    "routers",
					Aliases: []string{"r"},
					Usage:   "Routers dir",
				},
				&cli.StringFlag{
					Name:    "out",
					Aliases: []string{"o"},
					Usage:   "Routers file generate",
				},
			},
			Action: func(c *cli.Context) error {
				generate.GenerateRouter(c.String("routers"), c.String("controllers"), c.String("out"))
				return nil
			},
		},
	}
}