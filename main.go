package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/vindecodex/msgo/config"
	"github.com/vindecodex/msgo/controller"
	"github.com/vindecodex/msgo/domain"
	"github.com/vindecodex/msgo/logger"
	"github.com/vindecodex/msgo/middleware"
	"github.com/vindecodex/msgo/service"
)

func main() {
	logger.Info("MSGO starting...")
	router := mux.NewRouter()

	bookController := controller.BookController{
		service.NewDefaultBookService(domain.NewBookRepositoryAdapter(dbClient())),
	}
	userController := controller.UserController{
		service.NewDefaultUserService(domain.NewUserRepositoryAdapter(dbClient())),
	}

	authMiddleware := middleware.AuthMiddleware{domain.NewUserRepositoryAdapter(dbClient())}

	router.HandleFunc("/", controller.Welcome)

	router.HandleFunc("/books", bookController.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{book_id:[0-9]+}", bookController.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", bookController.NewBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{book_id:[0-9]+}", bookController.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{book_id:[0-9]+}", bookController.DeleteBook).Methods(http.MethodDelete)

	router.HandleFunc("/login", userController.Login).Methods(http.MethodPost)
	router.HandleFunc("/register", userController.Register).Methods(http.MethodPost)

	router.Use(authMiddleware.Authorize)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.GETSTRING("PORT")), router))
}

func dbClient() *sqlx.DB {
	dbUser := config.GETSTRING("DB_USER")
	dbPwd := config.GETSTRING("DB_PWD")
	dbHost := config.GETSTRING("DB_HOST")
	dbName := config.GETSTRING("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPwd, dbHost, dbName)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client

}
