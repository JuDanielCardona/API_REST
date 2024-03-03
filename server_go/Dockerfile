# Usa la imagen de Go como base
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el binario de tu aplicación al contenedor
COPY . .

# Expone el puerto 80 en el contenedor
EXPOSE 8080

ENTRYPOINT [ "go" ]

# Define el comando de inicio
CMD ["run", "./main.go"]

# Crear imagen con cd server_go y docker build -t server_go .
