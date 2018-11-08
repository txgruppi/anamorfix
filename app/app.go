package app

import (
	"github.com/txgruppi/anamorfix/resize"
	"gopkg.in/urfave/cli.v2"
)

func NewApp() *cli.App {
	app := cli.App{
		Name:      "anamorfix",
		HelpName:  "anamorfix",
		Usage:     "Fix aspect ratio of anamorphic images",
		ArgsUsage: "<files>",
		Version:   "0.0.0",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Tarcisio Gruppi",
				Email: "txgruppi@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "width",
				Aliases:     []string{"w"},
				Usage:       "Resize image's width",
				DefaultText: "resize image's height",
			},
			&cli.Float64Flag{
				Name:    "ratio",
				Aliases: []string{"r"},
				Usage:   "resize ratio",
				Value:   1.33,
			},
			&cli.StringFlag{
				Name:    "suffix",
				Aliases: []string{"s"},
				Usage:   "name suffix to add to resized files",
				Value:   "-desqueezed",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				cli.ShowAppHelpAndExit(c, 1)
				return nil
			}

			suffix := c.String("suffix")
			ratio := c.Float64("ratio")
			resizeWidth := c.Bool("width")

			config := resize.NewConfigWithDefaultEncoders(suffix, ratio, resizeWidth)

			return resize.All(config, c.Args().Slice())
		},
	}

	return &app
}
