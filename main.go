package main

import (
	"context"
	"latihanFSE/app"
	"latihanFSE/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.InitConfig()

	mysqlConn, sqlDB, errMysql := config.ConnectMySQL()
	if errMysql != nil {
		log.Println("error mysql connection: ", errMysql)
	}
	defer sqlDB.Close()

	hbaseConn, hbaseDB, errHBase := config.ConnectHbase()
	if errHBase != nil {
		log.Println("error HBase connection: ", errHBase)
	}
	defer hbaseDB.Close()

	if errMysql == nil {
		router := app.InitRouter(mysqlConn, hbaseConn)
		log.Println("routes Initialized")

		port := config.CONFIG["PORT"]
		srv := &http.Server{
			Addr:    ":" + port,
			Handler: router,
		}
		env := config.ENVIRONMENT
		log.Println("Server Initialized at " + env + ":" + port)

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		log.Println("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		log.Println("Server exiting")
	}
}
