package main

import "github.com/dexthrottle/trfine/internal/app"

const (
	dbName = "rp"
)

func main() {
	app.Run(dbName)
}
