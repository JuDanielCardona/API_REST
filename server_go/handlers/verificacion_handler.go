package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func Verificacion_handler(w http.ResponseWriter, r *http.Request) {

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

	// Verificar si la solicitud contiene la cabecera Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Error 401: No autorizado - Se requiere token JWT", http.StatusUnauthorized)
		fmt.Println("Solicitud http: " + r.URL.Path + " Send: Error 401")
		return
	}

	// Extraer el token JWT del encabezado Authorization
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	// Verificar la validez del token

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar el algoritmo de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de firma inesperado: %v", token.Header["alg"])
		}
		return []byte("12345"), nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Error 401: No autorizado - Token JWT inválido o vencido", http.StatusUnauthorized)
		fmt.Println("Solicitud HTTP:", r.URL.Path, "Send: Error 401 - Token JWT inválido")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["iss"] != "ingesis.uniquindio.edu.co" {
		http.Error(w, "Error 401: No autorizado - El emisor en el token no coincide", http.StatusUnauthorized)
		fmt.Println("Solicitud HTTP:", r.URL.Path, "Send: Error 401 - El emisor en el token no coincide")
		return
	}

	if claims["sub"] != username {
		http.Error(w, "Error 401: No autorizado - El nombre en el token no coincide", http.StatusUnauthorized)
		fmt.Println("Solicitud HTTP:", r.URL.Path, "Send: Error 401 - El nombre en el token no coincide")
		return
	}

	// Si todo está bien, responder con el saludo
	response := fmt.Sprintf("Hola, %s", username)
	fmt.Println("Solicitud HTTP:", r.URL.Path, "Send: 200 OK")
	fmt.Fprintln(w, response)

}
