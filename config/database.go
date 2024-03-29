package config

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "127.0.0.1"
	port     = 1433
	user     = "sa"
	password = "MK3ultimate"
	dbName   = "INVENTCONTROLGO"
)

func DatabaseConnection() *gorm.DB {

	// "sqlserver://seu_usuario:senha@localhost:1433?database=seu_banco_de_dados"
	sqlInfo := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, password, host, port, dbName)
	db, err := gorm.Open(sqlserver.Open(sqlInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Falha ao conectar ao banco de dados: " + err.Error())
	}

	return db
}
