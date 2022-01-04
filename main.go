package main

import "mongodb-course/delivery"

func main() {
	var server delivery.Routes
	server.StartGin()
}
