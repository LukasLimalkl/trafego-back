package main

import (
	"github.com/LukasLimalkl/trafego-back/internal/router"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	router.Initialize()

}
