package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/vindecodex/msgo/config"
	"github.com/vindecodex/msgo/logger"
)

var (
	dbUser = config.GETSTRING("DB_USER")
	dbPwd  = config.GETSTRING("DB_PWD")
	dbHost = config.GETSTRING("DB_HOST")
	dbName = config.GETSTRING("DB_NAME")
)

var schemaBooks = `
CREATE TABLE IF NOT EXISTS books (
	id int(11) NOT NULL AUTO_INCREMENT,
	title varchar(90),
	author varchar(60),
	length int(11), 
	PRIMARY KEY (id)
);`

var schemaUsers = `
CREATE TABLE IF NOT EXISTS users (
id int(11) NOT NULL AUTO_INCREMENT,
username varchar(40) UNIQUE,
password varchar(40),
role varchar(10),
PRIMARY KEY (id)
);`

func initializeDatabase() {
	logger.Info("initializeDatabase")
	client, err := dbClient()
	if err != nil {
		log.Printf("Error %s when opening database\n", err)
		return
	}
	client.Exec("CREATE DATABASE IF NOT EXISTS msgo")
	tx := client.MustBegin()
	client.MustExec(schemaBooks)
	client.MustExec(schemaUsers)

	tx.MustExec("INSERT IGNORE INTO users (username, password, role) VALUES (\"admin\", \"admin\", \"owner\")")
	tx.Commit()
}

func dbClient() (*sqlx.DB, error) {
	logger.Info("dbClient")
	dataSource := getDataSource("")
	client, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	client.Exec("CREATE DATABASE IF NOT EXISTS msgo")

	client, err = sqlx.Open("mysql", getDataSource(dbName))
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client, nil

}

func getDataSource(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPwd, dbHost, dbname)
}
