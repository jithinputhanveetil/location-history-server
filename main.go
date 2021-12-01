package main

import "location-history-server/internal/app"

func main() {
	s := app.NewServer()
	s.Start()
}
