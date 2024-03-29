package base

import (
	"log"
	"serveit/tools"
)

type Option func(app *Application)

func SetProfile(profileFilePath string) Option {
	return func(app *Application) {
		var err error
		app.profile, err = tools.ParseToml(profileFilePath)
		if err != nil {
			log.Printf("[ERROR] parse profile failed, %v", err)
		}
	}
}
