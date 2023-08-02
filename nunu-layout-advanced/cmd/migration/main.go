package main

import (
	"nunu-layout-advanced/pkg/config"
	"nunu-layout-advanced/pkg/log"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	app, cleanup, err := newApp(conf, logger)
	if err != nil {
		panic(err)
	}
	app.Run()
	defer cleanup()
}
