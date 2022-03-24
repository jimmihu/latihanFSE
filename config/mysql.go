package config

import (
	"database/sql"
	"log"

	"latihanFSE/models/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL() (*gorm.DB, *sql.DB, error) {

	connectionString := CONFIG["MYSQL_USER"] + ":" + CONFIG["MYSQL_PASS"] + "@tcp(" + CONFIG["MYSQL_HOST"] + ":" + CONFIG["MYSQL_PORT"] + ")/" + CONFIG["MYSQL_SCHEMA"] + "?parseTime=true"
	mysqlConn, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Println("Error connect to MySQL: ", err.Error())
		return nil, nil, err
	}

	sqlDB, errDB := mysqlConn.DB()
	if errDB != nil {
		log.Println(errDB)
	} else {
		sqlDB.SetMaxIdleConns(2)
		sqlDB.SetMaxOpenConns(1000)
	}
	mysqlConn.AutoMigrate(&entity.User{})
	mysqlConn.AutoMigrate(&entity.Role{})

	//seed roles
	role := entity.Role{}
	var roles = [5]string{"admin", "maker", "checker", "signer", "viewer"}
	if err := mysqlConn.Where("title = ?", "admin").First(&role).Error; err != nil {
		for i := 0; i < len(roles); i++ {
			mysqlConn.Create(&entity.Role{Title: roles[i], Active: true})
		}
	}

	log.Println("MySQL connection success")
	return mysqlConn, sqlDB, nil

	// return nil, nil
}
