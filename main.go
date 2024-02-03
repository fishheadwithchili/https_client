package main

import (
	"https_client/component"
)

func main() {
	component.LoadConfig("config.yml")
	var client = component.NewClient()
	client.Run()
}
