package main

import (
	"fmt"

	"github.com/isbang/dsn/postgres"
	"gopkg.in/AlecAivazis/survey.v1"
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

	var cfg postgres.Config
	if err := survey.Ask(question, &cfg); err != nil {
		panic(err)
	}
	fmt.Println(cfg.DSN())
}
