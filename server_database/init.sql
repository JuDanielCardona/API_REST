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
('Juan', 'perro', 'juan@email.com'),
('Pedro', 'gato', 'pedro@email.com'),
('Luis', 'casa', 'luis@email.com'),
('Diego', 'sol', 'diego@email.com'),
('Carlos', 'playa', 'carlos@email.com'),
('Sofía', 'luna', 'sofia@email.com'),
('Ana', 'mar', 'ana@email.com'),
('María', 'montaña', 'maria@email.com'),
('Lucía', 'agua', 'lucia@email.com'),
('Miguel', 'aire', 'miguel@email.com'),
('Andrés', 'fuego', 'andres@email.com'),
('Santiago', 'tierra', 'santiago@email.com'),
('Camila', 'árbol', 'camila@email.com'),
('Valentina', 'cielo', 'valentina@email.com'),
('Juliana', 'nube', 'juliana@email.com'),
('Antonio', 'lluvia', 'antonio@email.com'),
('Javier', 'viento', 'javier@email.com'),
('Gabriel', 'nieve', 'gabriel@email.com'),
('Mateo', 'rayo', 'mateo@email.com'),
('Daniel', 'trueno', 'daniel@email.com'),
('Lorenzo', 'rayo', 'lorenzo@email.com'),
('Manuel', 'trueno', 'manuel@email.com'),
('Ismael', 'lluvia', 'ismael@email.com'),
('Raúl', 'viento', 'raul@email.com'),
('Joaquín', 'nieve', 'joaquin@email.com'),
('Fernando', 'nube', 'fernando@email.com'),
('Adrián', 'cielo', 'adrian@email.com'),
('Félix', 'árbol', 'felix@email.com'),
('Rafael', 'tierra', 'rafael@email.com');



--psql -U admin -d users_db -f init.sql