package main

import (
	"bss-beta-backend/app"
	"log"
)

func main() {
	// config.SetDefault(data.MigrationPathConfig, "./api")
	// config.SetDefault(data.DatabaseURLConfig, "postgres:///postgres?sslmode=disable")
	err := app.Build().Start()
	if err != nil {
		log.Fatalln(err)
	}
}
