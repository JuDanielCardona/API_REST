# Usa una imagen base de MySQL
FROM mysql:latest

# Define la variable de entorno MYSQL_ROOT_PASSWORD
ENV MYSQL_ROOT_PASSWORD 12345

# Copia el archivo init.sql al contenedor
COPY init.sql /docker-entrypoint-initdb.d/

# Cambia el directorio de trabajo al directorio de scripts de inicialización
WORKDIR /docker-entrypoint-initdb.d/

#docker run --name server_db server_database