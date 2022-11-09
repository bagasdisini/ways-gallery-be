package postdto

type CreatePostRequest struct {
	UserID int    `json:"userID"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Date   string `json:"date"`
	Post1  string `json:"post1"`
	Post2  string `json:"post2"`
	Post3  string `json:"post3"`
	Post4  string `json:"post4"`
	Post5  string `json:"post5"`
}

type PostResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"title" `
	Desc   string `json:"desc"`
	Date   string `json:"date"`
	UserID int    `json:"userID"`
	Post1  string `json:"post1" `
	Post2  string `json:"post2" `
	Post3  string `json:"post3" `
	Post4  string `json:"post4" `
	Post5  string `json:"post5" `
}
