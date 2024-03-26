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
	login_EndPoints(router.PathPrefix("/apiSesion").Subrouter())
	funtions_EndPoints(router.PathPrefix("/apiUser").Subrouter())

	fmt.Println("Init server on port: " + PORT)
	http.ListenAndServe(PORT, router) // Aquí se pasa el router en lugar de nil
}

func login_EndPoints(router *mux.Router) {
	//http://localhost:8080/apiSesion/
	router.HandleFunc("/login", handlers.Login_handler).Methods("POST")
	router.HandleFunc("/recover/{email}", handlers.RecoverPassword_handler).Methods("GET")
	router.HandleFunc("/update", handlers.UpdatePassword_handler).Methods("POST")
}

func funtions_EndPoints(router *mux.Router) {
	//http://localhost:8080/apiUser/
	router.HandleFunc("/add", handlers.AddUser_handler).Methods("POST")
	router.HandleFunc("/update", handlers.UpdateUser_handler).Methods("PUT")
	router.HandleFunc("/delete", handlers.DeleteUser_handler).Methods("DELETE")
	router.HandleFunc("/search/{id}", handlers.GetUserById_handler).Methods("GET")
	router.HandleFunc("/all", handlers.GetAllUsers_handler).Methods("GET")
}
