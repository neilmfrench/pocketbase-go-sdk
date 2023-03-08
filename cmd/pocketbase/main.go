package main

import (
	_ "github.com/neilmfrench/pocketbase-go-sdk"
	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		panic(err)
	}
}
