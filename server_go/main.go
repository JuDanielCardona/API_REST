package main

import (
	"fmt"
	"net/http"
	"taller_docker/database"
	"taller_docker/handlers"
	"taller_docker/models"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("-----------------INIT SERVER GO-----------------")
	database.Connection()
	database.DB.AutoMigrate(models.User{})
	router := mux.NewRouter()
	PORT := ":8080"
	login_EndPoints(router.PathPrefix("/apiLogin").Subrouter())
	funtions_EndPoints(router.PathPrefix("/apiUser").Subrouter())

	fmt.Println("Init server on port: " + PORT)
	http.ListenAndServe(PORT, router) // Aquí se pasa el router en lugar de nil
}

func login_EndPoints(router *mux.Router) {
	//http://localhost:8080/apiLogin/login
	router.HandleFunc("/login", handlers.Login_handler).Methods("POST")
}

func funtions_EndPoints(router *mux.Router) {
	//http://localhost:8080/apiUser/
	router.HandleFunc("/add", handlers.AddUser_handler).Methods("POST")
	router.HandleFunc("/all", handlers.GetAllUsers_handler).Methods("GET")
	router.HandleFunc("/{id}", handlers.GetUserById_handler).Methods("GET")
	router.HandleFunc("/delete", handlers.DeleteUser_handler).Methods("DELETE")
	router.HandleFunc("/{id}", handlers.UpdateUser_handler).Methods("PUT")
}
