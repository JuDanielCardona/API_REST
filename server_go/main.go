package main

import (
	"fmt"
	"net/http"
	"taller_docker/database"
	//	"taller_docker/handlers"
)

func main() {
	database.Connection()
	//http.HandleFunc("/saludo", handlers.Saludo_handler)
	//http.HandleFunc("/saludo", handlers.Verificacion_handler)
	//http.HandleFunc("/login", handlers.Login_handler)

	fmt.Println("Init server on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
