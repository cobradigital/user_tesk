package repositories

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(host, port, user, pass, name string) (*gorm.DB, error) {
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
