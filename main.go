package main

import (
	"fmt"
	"sg/config"
)

func main() {
	c := config.GetConfig()
	config.GetConfig()

	fmt.Println(c.GetModo())
}
