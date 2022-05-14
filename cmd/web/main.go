package main

import "github.com/dexthrottle/trfine/internal/app"

const (
	dbName  = "rp"
	baseURL = "https://api-testnet.bybit.com/"
)

func main() {
	app.Run(dbName, baseURL)
}
