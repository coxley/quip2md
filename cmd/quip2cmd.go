package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/coxley/quip2md"
	"github.com/urfave/cli"
)

func convert(c *cli.Context) error {
	fileName := c.Args().Get(0)

	var data []byte
	var err error
	if fileName == "" {
		data, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
	} else {
		data, err = ioutil.ReadFile(fileName)
		if err != nil {
			return err
		}
	}
	html := string(data)
	converted, err := quip2md.QuipToMarkdown(html)
	if err != nil {
		return err
	}
	fmt.Printf(converted)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Action = convert
	app.Name = "quip2md"
	app.Usage = "Convert Quip HTML to Markdown"
	app.UsageText = "quip2md [file to convert]\n   cat [file to convert] | quip2md\n   custom_program | quip2md"
	app.Version = "1.0.1"

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
