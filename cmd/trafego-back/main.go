package main

import (
	"fmt"

	"github.com/LukasLimalkl/trafego-back/config"
	"github.com/LukasLimalkl/trafego-back/internal/router"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	// Initialize Configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize Routes
	router.Initialize()

}
