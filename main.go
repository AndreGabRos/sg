package main

import (
	"fmt"
	"sg/config"
)

func main() {
	c := config.GetConfig()
	c2 := config.GetConfig()

	fmt.Println(c.GetModo())
	c.SetModo("oi")
	fmt.Println(c2.GetModo())
}
