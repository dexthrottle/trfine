package main

import "trfine/internal/app"

const (
	saveToFile = false
	baseURL    = "https://api-testnet.bybit.com/"
	tgBotDebug = false
)

func main() {
	app.RunApplication(saveToFile)
}
