package main

import (
	"os"

	"gopkg.in/pot.v1/cli"
	"gopkg.in/pot.v1/cmd/pot/generate"
	"gopkg.in/pot.v1/cmd/pot/logo"
	"gopkg.in/pot.v1/cmd/pot/run"
)

func main() {

	app := cli.App{}

	app.Name = "pot"

	app.Usage = "Pot is web framework for the Go programming language."

	app.Version = version

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
			logo.PrintLogo("V" + version)
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
					Name:    "package",
					Aliases: []string{"p"},
					Usage:   "Controllers dir",
				},
				&cli.StringFlag{
					Name:    "out",
					Aliases: []string{"o"},
					Usage:   "Routers file generate",
				},
			},
			Action: func(c *cli.Context) error {
				generate.GenerateRouter(c.String("package"), c.String("routers"), c.String("controllers"), c.String("out"))
				return nil
			},
		},
	}
}
