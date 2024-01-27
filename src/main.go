package main

import "backend-golang-gin/app"

func main() {
	var app app.App

	app.Connect()

	app.CreateRoutes()

	app.Run()
}
