package main

import (
	"fmt"

	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/isbang/dsn/mysql"
)

func main() {
	question := []*survey.Question{
		{
			Name:   "host",
			Prompt: &survey.Input{Message: "host"},
		}, {
			Name:   "port",
			Prompt: &survey.Input{Message: "port"},
		}, {
			Name:   "user",
			Prompt: &survey.Input{Message: "user"},
		}, {
			Name:   "pass",
			Prompt: &survey.Input{Message: "pass"},
		}, {
			Name:   "db",
			Prompt: &survey.Input{Message: "db"},
		},
	}

	var cfg mysql.Config
	if err := survey.Ask(question, &cfg); err != nil {
		panic(err)
	}
	fmt.Println(cfg.DSN())
}
