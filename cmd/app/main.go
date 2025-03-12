package main

import "music/internal/app"

func main() {
	_, err := app.NewApp()
	if err != nil {
		panic(err)
	}
}
