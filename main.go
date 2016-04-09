package main

import (
	"github.com/AntonYaskevich/lu-server/router"
)

const (
	Port = "8080"
)

func main() {
	engine := router.CreateRouterEngine()
	engine.Run(":" + Port)
}