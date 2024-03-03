package handlers

import (
	"fmt"
	"net/http"
)

// Poner nombre Saludo_handler para probar punto 1
func Saludo_handler(w http.ResponseWriter, r *http.Request) {

	username := r.URL.Query().Get("nombre")

	if r.URL.Path != "/saludo" {
		http.Error(w, "Error 404: Recurso no encontrado", http.StatusNotFound)
		return
	}

	if username == "" {
		http.Error(w, "Error 400: Solicitud no valida - el nombre es obligatorio", http.StatusBadRequest)
		fmt.Println("Solicitud http: " + r.URL.Path + " Send: Error 400")
		return
	}

	response := fmt.Sprintf("Hola, %s", username)
	fmt.Println("Solicitud http: " + r.URL.Path + " Send: 200 ok")
	fmt.Fprintln(w, response)

}
