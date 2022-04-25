package main

import "github.com/dexthrottle/trfine/internal/app"

const (
	appPort = "8000"
	ginMode = "debug"
)

func main() {
	app.Run(appPort, ginMode)
}
