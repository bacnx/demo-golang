package main

import (
	"go-gin/config"
	"go-gin/router"
)

func main() {
	config.SetEnv()
	app := router.NewService()
	if err := app.Start(); err != nil {
		panic(err)
	}
}
