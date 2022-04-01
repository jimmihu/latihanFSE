package config

import (
	"database/sql"
	"log"

	_ "github.com/apache/calcite-avatica-go/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectHbase() (*gorm.DB, *sql.DB, error) {

	connectionString := CONFIG["PQS_URL"]
	avatica, err := sql.Open("avatica", connectionString)
	hbaseConn, err := gorm.Open(mysql.New(mysql.Config{Conn: avatica}), &gorm.Config{})
	if err != nil {
		log.Println("DB connection error: ", err.Error())
	}
	hbaseDB, errDB := hbaseConn.DB()
	if errDB != nil {
		log.Println(errDB)
	} else {
		hbaseDB.SetMaxIdleConns(2)
		hbaseDB.SetMaxOpenConns(1000)
	}

	migrate := hbaseConn.Raw("CREATE TABLE IF NOT EXISTS ?", "users")
	if migrate.Error != nil {
		log.Println("HBase migration error: ", migrate.Error.Error())
	}

	log.Println("Hbase connnection success: ", CONFIG["PQS_URL"])
	return hbaseConn, hbaseDB, nil
}
