package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Credentials struct {
	Usuario string `json:"usuario"`
	Clave   string `json:"clave"`
}

func main() {

	SERVER_URL := os.Getenv("SERVER_URL")
	User := generate(10)
	Pass := generate(32)
	fmt.Println("\nUser: " + User +
		"\nPass: " + Pass +
		"\nRequests to " + SERVER_URL)

	// Generar credenciales aleatorias
	credenciales := Credentials{
		Usuario: User,
		Clave:   Pass,
	}

	// Convertir credenciales a formato JSON
	credencialesJSON, err := json.Marshal(credenciales)
	if err != nil {
		log.Fatal("\nError: No se convirtieron las credenciales", err)
	} else {
		log.Println("Credentials it´s OK generated")
	}

	// Realizar solicitud POST a la ruta de login del server_go
	respLogin, err := http.Post(SERVER_URL+"/login", "application/json", bytes.NewBuffer(credencialesJSON))
	if err != nil {
		log.Fatal("\nError: No se pudo hacer la solicitud de login", err)
		return
	} else {
		log.Println("Login it´s OK required")
	}
	defer respLogin.Body.Close()

	// Leer la respuesta del servidor
	respLoginBody, err := ioutil.ReadAll(respLogin.Body)
	if err != nil {
		log.Fatal("\nError: No se recibió respuesta del servidor", err)
		return
	} else {
		log.Println("Server response: " + string(respLoginBody))
	}

	// Extraer el token de la respuesta del servidor
	token := string(respLoginBody)

	// Realizar solicitud GET a la ruta de saludo del server_go
	req, err := http.NewRequest("GET", SERVER_URL+"/saludo?nombre="+User, nil)
	if err != nil {
		fmt.Println("Error al crear la solicitud de saludo:", err)
		return
	}

	// Agregar el token JWT al encabezado de autorización
	req.Header.Set("Authorization", "Bearer "+token)

	// Realizar la solicitud GET con el token JWT en el encabezado de autorización
	respSaludo, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("\nError: No se pudo hacer la solicitud de saludo", err)
		return
	} else {
		log.Println("Saludo it´s OK required")
	}
	defer respSaludo.Body.Close()

	// Leer la respuesta del servidor
	respSaludoBody, err := ioutil.ReadAll(respSaludo.Body)
	if err != nil {
		log.Fatal("\nError: No se recibió respuesta del servidor", err)
		return
	} else {
		log.Println("Server response: " + string(respSaludoBody))
	}

}

func generate(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// La fuente global rand debe estar inicializada
	rand.Seed(time.Now().UnixNano())

	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}
