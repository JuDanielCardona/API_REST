package database

import (
	"errors"
	"fmt"
	"taller_docker/models"

	"gorm.io/gorm"
)

func SearchUser(user *models.User) (bool, error) {
	if err := DB.Where("name = ? AND password = ?", user.Name, user.Password).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func GetUsers(page, pageSize int) ([]models.User, error) {
	var users []models.User

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
	if err := DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Error: User not found")
		}
		return nil, err
	}
	return &user, nil
}

func AddUser(addUser models.User) (*models.User, error) {

	var existingUser models.User
	if err := DB.Where("email = ?", addUser.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("Error")
	}

	newUser := DB.Create(&addUser)
	err := newUser.Error

	if err != nil {
		return nil, err
	}
	return &addUser, nil
}

func DeleteUser(deletedUser models.User) error {
	var user models.User
	if err := DB.Where("id = ?", deletedUser.Id).First(&user).Error; err != nil {
		return errors.New("user not found")
	}

	if err := DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(updatedUser models.User) (models.User, error) {
	var user models.User
	if err := DB.Where("id = ?", updatedUser.Id).First(&user).Error; err != nil {
		return models.User{}, errors.New("user not found")
	}

	oldPassword := user.Password

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Date = updatedUser.Date
	if user.Password != updatedUser.Password {
		fmt.Println("LOG: Password change prevented.")
	}

	if err := DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}

	if err := DB.Where("id = ?", updatedUser.Id).First(&user).Error; err != nil {
		return models.User{}, errors.New("failed to reload user from database")
	}

	user.Password = oldPassword

	return user, nil
}

func RecoverPassword(email string) (string, error) {
	var userToUpdate models.User
	DB.Where("email = ?", email).First(&userToUpdate)

	if userToUpdate.Password == "" {
		return "", errors.New("Error: User not found")
	}
	return userToUpdate.Password, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdatePassword(user models.User, email string) (string, error) {
	var userToUpdate models.User
	if err := DB.Where("email = ?", email).First(&userToUpdate).Error; err != nil {
		return "", errors.New("Error: User not found")
	}

	// Verificar que la nueva contraseña sea diferente de la anterior
	if userToUpdate.Password == user.Password {
		return "", errors.New("Error: New password must be different from the current one")
	}

	// Actualizar la contraseña del usuario
	userToUpdate.Password = user.Password
	if err := DB.Save(&userToUpdate).Error; err != nil {
		return "", err
	}

	return userToUpdate.Password, nil
}
