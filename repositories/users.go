package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	ShowUsers() ([]models.User, error)
	GetUserByIDUser(ID int) (models.User, error)
	CreateUserUser(user models.User) (models.User, error)
	UpdateUser(user models.User, ID int) (models.User, error)
	DeleteUser(user models.User, ID int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ShowUsers() ([]models.User, error) {
	var user []models.User
	err := r.db.Find(&user).Error

	return user, err
}

func (r *repository) GetUsers(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}

func (r *repository) CreateUserUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
	err := r.db.Model(&user).Where("id=?", ID).Updates(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User, ID int) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
