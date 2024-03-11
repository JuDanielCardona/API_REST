-- Elimina la base de datos
DROP DATABASE IF EXISTS users_db;
-- Crear la base de datos si no existe
CREATE DATABASE users_db;

-- Conectar a la base de datos recién creada
\c users_db;

-- Crear la tabla de usuarios si no existe
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insertar datos de prueba en la tabla users
INSERT INTO users (name, password, email) VALUES
('Juan', 'contrasena', 'juanperez@gmail.com'),
('Maria', 'contrasena', 'mariagarcia@gmail.com'),
('Luis', 'contrasena', 'luisrodriguez@gmail.com'),
('Ana', 'contrasena', 'anamartinez@gmail.com'),
('Jose', 'contrasena', 'joselopez@gmail.com'),
('Sofia', 'contrasena', 'sofiahernandez@gmail.com'),
('Carlos', 'contrasena', 'carlosgomez@gmail.com'),
('Laura', 'contrasena', 'lauradiaz@gmail.com'),
('Pedro', 'contrasena', 'pedroruiz@gmail.com'),
('Monica', 'contrasena', 'monicasanchez@gmail.com');


--psql -U admin -d users_db -f init.sql