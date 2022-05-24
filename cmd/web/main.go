package main

import "github.com/dexthrottle/trfine/internal/app"

const (
	dbName     = "rp"
	baseURL    = "https://api-testnet.bybit.com/"
	tgBotDebug = false
)

func main() {
	app.Run(dbName, baseURL, tgBotDebug)
}
