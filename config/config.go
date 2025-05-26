package config

import (
	"fmt"
	"os"
	"sync"
)

type config struct {
	modo string
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

		instance = &config{}

		instance.modo = modo
	})

	return instance
}

func (c config) GetModo() string {
	return c.modo
}

func (c *config) SetModo(modo string) {
	c.modo = modo
}
