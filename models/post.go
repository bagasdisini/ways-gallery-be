package models

type Post struct {
	ID     int    `json:"id" gorm:"primary_key:auto_increment;"`
	Title  string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Desc   string `json:"desc" form:"desc" gorm:"type: varchar(255)"`
	Date   string `json:"date" form:"date" gorm:"type: varchar(255)"`
	UserID int    `json:"userID"`
	Post1  string `json:"post1" form:"post1" gorm:"type: varchar(255)"`
	Post2  string `json:"post2" form:"post2" gorm:"type: varchar(255)"`
	Post3  string `json:"post3" form:"post3" gorm:"type: varchar(255)"`
	Post4  string `json:"post4" form:"post4" gorm:"type: varchar(255)"`
	Post5  string `json:"post5" form:"post5" gorm:"type: varchar(255)"`
}
