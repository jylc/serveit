package base

type Option func(app *Application)

func SetProfile(profileFilePath string) Option {
	return func(app *Application) {
		app.profile = profileFilePath
	}
}
