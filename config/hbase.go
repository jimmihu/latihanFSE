package config

import (
	"errors"
	"log"
	"strings"

	_ "github.com/apache/calcite-avatica-go/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectHbase() ([]*gorm.DB, error) {

	listUrl := strings.Split(CONFIG["PQS_URL"], ",")

	var listConn []*gorm.DB
	for _, url := range listUrl {
		hbaseConn, err := gorm.Open("avatica", url)
		if err != nil {
			log.Println("DB connection error: ", err.Error())
			continue
		}
		hbaseConn.DB().SetMaxIdleConns(2)
		hbaseConn.DB().SetMaxOpenConns(1000)
		hbaseConn.LogMode(true)
		log.Println("Hbase connnection success: ", url)
		listConn = append(listConn, hbaseConn)
	}

	if len(listConn) == 0 {
		return nil, errors.New("no hbase connection available")
	}

	return listConn, nil
}
