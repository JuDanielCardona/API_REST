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
('Juan Pérez', 'contraseña1', 'juanperez@gmail.com'),
('María García', 'contraseña2', 'mariagarcia@gmail.com'),
('Luis Rodríguez', 'contraseña3', 'luisrodriguez@gmail.com'),
('Ana Martínez', 'contraseña4', 'anamartinez@gmail.com'),
('José López', 'contraseña5', 'joselopez@gmail.com'),
('Sofía Hernández', 'contraseña6', 'sofiahernandez@gmail.com'),
('Carlos Gómez', 'contraseña7', 'carlosgomez@gmail.com'),
('Laura Díaz', 'contraseña8', 'lauradiaz@gmail.com'),
('Pedro Ruiz', 'contraseña9', 'pedroruiz@gmail.com'),
('Mónica Sánchez', 'contraseña10', 'monicasanchez@gmail.com');


--psql -U admin -d users_db -f init.sql