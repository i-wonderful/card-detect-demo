package main

import (
	. "card-detect-demo/internal/app"
	"log"
)

func main() {
	config, err := NewConfigFromYml()
	if err != nil {
		log.Fatal(err)
		return
	}
	app := NewApp(config)
	log.Fatal(app.Run())
}
