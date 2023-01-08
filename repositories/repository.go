package repositories

import (
	"gorm.io/gorm"
)

// Untuk di gunakan secara global
type repository struct {
	db *gorm.DB
}
