package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/vindecodex/msgo/config"
	"github.com/vindecodex/msgo/controller"
	"github.com/vindecodex/msgo/domain"
	"github.com/vindecodex/msgo/logger"
	"github.com/vindecodex/msgo/middleware"
	"github.com/vindecodex/msgo/service"
)

var client *sqlx.DB

func init() {
	initializeDatabase()
	dbclient, err := dbClient()
	if err != nil {
		logger.Error("Error on initialization of Database")
		panic(err)
	}
	client = dbclient
}

func main() {
	logger.Info("MSGO starting...")
	router := mux.NewRouter()

	bookController := controller.BookController{
		service.NewDefaultBookService(domain.NewBookRepositoryAdapter(client)),
	}
	userController := controller.UserController{
		service.NewDefaultUserService(domain.NewUserRepositoryAdapter(client)),
	}

	authMiddleware := middleware.AuthMiddleware{domain.NewUserRepositoryAdapter(client)}

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
