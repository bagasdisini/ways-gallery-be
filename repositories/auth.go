package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterUser(user models.User) (models.User, error)
	LoginUser(email string) (models.User, error)
	GetUsers(ID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RegisterUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) LoginUser(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) GetUserByIDUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}
