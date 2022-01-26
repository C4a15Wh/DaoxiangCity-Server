package db

import (
	"dxcserver/model"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var exeDB *gorm.DB

func OpenDB(config model.ServerConf) {
	var err error

	ConnectInfo := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		config.SQL.User, config.SQL.Password, config.SQL.Addr, config.SQL.Port, config.SQL.DB)

	exeDB, err = gorm.Open("mysql", ConnectInfo)

	if err != nil {
		log.Fatalf("Failed to connect to DB: %s", err)
	}
}

func CloseDB() {
	if exeDB != nil {
		_ = exeDB.Close()
	}
}
