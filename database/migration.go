package database

import (
	"backend/models"
	"backend/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Post{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed!")
	}

	fmt.Println("Migration Successful!")
}
