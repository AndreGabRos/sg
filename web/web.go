package web

import (
	"fmt"
	"sg/config"
)

func RodarServidorWeb() {
	c := config.GetConfig()

	fmt.Printf("Rodando servidor na porta %d no modo %s\n", c.GetPorta(), c.GetModo())
}
