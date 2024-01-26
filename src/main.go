package main

import "backend-golang-gin/tasks/infrastructure"

func main() {
	var app infrastructure.App
	app.SendMessage()
	app.Run()
}
