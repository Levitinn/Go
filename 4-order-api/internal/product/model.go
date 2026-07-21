package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `gorm:"not null"`
	Description string         `gorm:"not null"`
	Images      pq.StringArray `gorm:"type:text[]"`
}
