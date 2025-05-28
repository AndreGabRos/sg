package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type config struct {
	modo  string
	porta int
}

var instance *config
var once sync.Once

func GetConfig() *config {
	once.Do(func() {
		fmt.Println("Dentro do once")
		modo := os.Getenv("MODO")
		if modo == "" {
			modo = "desenvolvimento"
		}
		porta, err := strconv.Atoi(os.Getenv("PORTA"))
		if err != nil {
			porta = 2222
		}

		instance = &config{
			modo:  modo,
			porta: porta,
		}
	})

	return instance
}

func (c config) GetModo() string {
	return c.modo
}

func (c *config) SetModo(modo string) {
	c.modo = modo
}

func (c config) GetPorta() int {
	return c.porta
}

func (c *config) SetPorta(porta int) {
	c.porta = porta
}
