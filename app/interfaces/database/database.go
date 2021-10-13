package database

import "gorm.io/gorm"

type Database interface {
	Gorm() *gorm.DB
}
