package database

import (
	"errors"
	"fmt"
	"taller_docker/models"
)

func SearchUser(user *models.User) (bool, error) {
	if err := DB.Where("name = ? AND password = ?", user.Name, user.Password).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	page := 1      // Número de página predeterminado
	pageSize := 10 // Tamaño de página predeterminado

	// Calcula el desplazamiento basado en la página y el tamaño de la página
	offset := (page - 1) * pageSize

	// Realiza la consulta con el desplazamiento y el tamaño de página adecuados
	err := DB.Offset(offset).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserById(id string) (*models.User, error) {
	var user models.User
	DB.First(&user, id)
	if user.Id == 0 {
		return nil, errors.New("Error: User not found")
	}
	return &user, nil
}

func AddUser(user models.User) (*models.User, error) {
	newUser := DB.Create(&user)
	err := newUser.Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(id int) error {
	fmt.Println("SE INGRESA ID: ", id)
	var user models.User
	if err := DB.First(&user, id).Error; err != nil {
		return errors.New("user not found")
	}

	if user.Id == 0 {
		return errors.New("user not found")
	}

	if err := DB.Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(id int, updatedUser models.User) error {
	fmt.Println("Se ingresó ID: ", id)
	var user models.User
	if err := DB.Where("id = ?", id).First(&user).Error; err != nil {
		return errors.New("user not found")
	}

	if user.Id == 0 {
		return errors.New("user not found")
	}

	// Actualizar los campos del usuario
	user.Name = updatedUser.Name
	user.Password = updatedUser.Password
	user.Email = updatedUser.Email
	user.Date = updatedUser.Date

	// Guardar los cambios en la base de datos
	if err := DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
