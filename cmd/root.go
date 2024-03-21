package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"serveit/base"
	"serveit/tools"
)

var (
	profileFilePath string
	loggerFilePath  string
)

var rootCmd = &cobra.Command{
	Use:   "serveit",
	Short: "serveit is a distributed server",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		profileFilePath, err = tools.JudgeFile(profileFilePath)
		if err != nil {
			log.Fatal("[ERROR] cannot open profile path")
		}

		loggerFilePath, err = tools.JudgeFile(loggerFilePath)
		if err != nil {
			log.Fatal("[ERROR] cannot open log path")
		}

		loggerFile, err := os.OpenFile(loggerFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("[ERROR] cannot open log file")
		}
		multiOutput := io.MultiWriter(loggerFile, os.Stdin)
		log.SetOutput(multiOutput)
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

		app := base.NewApplication(base.SetProfile(profileFilePath))
		err = app.Startup()
		if err != nil {
			log.Fatal("[ERROR] serveit startup failed")
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&profileFilePath, "profile", "", "./config/profile.toml", "basic running information")
	rootCmd.Flags().StringVarP(&loggerFilePath, "logger", "", "./log", "log file directory")
}
