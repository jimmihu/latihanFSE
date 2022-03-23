package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectMySQL() (*gorm.DB, error) {

	connectionString := CONFIG["MYSQL_USER"] + ":" + CONFIG["MYSQL_PASS"] + "@tcp(" + CONFIG["MYSQL_HOST"] + ":" + CONFIG["MYSQL_PORT"] + ")/" + CONFIG["MYSQL_SCHEMA"] + "?parseTime=true"
	mysqlConn, err := gorm.Open("mysql", connectionString)

	if err != nil {
		log.Println("Error connect to MySQL: ", err.Error())
		return nil, err
	}

	mysqlConn.DB().SetMaxIdleConns(2)
	mysqlConn.DB().SetMaxOpenConns(9999)
	mysqlConn.LogMode(true)

	log.Println("MySQL connection success")
	return mysqlConn, nil

	// return nil, nil
}
