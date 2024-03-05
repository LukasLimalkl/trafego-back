package schema

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Nome string
	User string
}
