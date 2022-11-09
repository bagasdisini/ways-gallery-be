package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	ShowPosts() ([]models.Post, error)
	GetPostByID(ID int) (models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
}

func RepositoryPost(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ShowPosts() ([]models.Post, error) {
	var Posts []models.Post
	err := r.db.Preload("User").Find(&Posts).Error

	return Posts, err
}

func (r *repository) GetPostByID(ID int) (models.Post, error) {
	var Post models.Post
	err := r.db.Preload("User").First(&Post, ID).Error

	return Post, err
}

func (r *repository) CreatePost(Post models.Post) (models.Post, error) {
	err := r.db.Create(&Post).Error

	return Post, err
}
