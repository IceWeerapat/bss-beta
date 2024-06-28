package home

import (
	"log"

	"fx.prodigy9.co/app"
	"github.com/spf13/cobra"
)

var HelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Logs hello world, you know, for testing your sanity.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hello, World!")
	},
}

var App = app.Build().
	Command(HelloCmd).
	Controllers(HomeCtr{})
