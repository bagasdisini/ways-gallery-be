package transactiondto

type CreateTransactionRequest struct {
	AdminID   int    `json:"admin_id"`
	BuyerID   int    `json:"buyer_id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Price     string `json:"price"`
}

type UpdateTransactionRequest struct {
	Status      string `json:"status"`
	ProjectDesc string `json:"projectDesc"`
	Project1    string `json:"project1"`
	Project2    string `json:"project2"`
	Project3    string `json:"project3"`
	Project4    string `json:"project4"`
	Project5    string `json:"project5"`
}

type TransactionResponse struct {
	ID          int    `json:"id"`
	AdminID     int    `json:"admin_id"`
	BuyerID     int    `json:"buyer_id"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Price       string `json:"price"`
	Status      string `json:"status"`
	ProjectDesc string `json:"projectDesc"`
	Project1    string `json:"project1"`
	Project2    string `json:"project2"`
	Project3    string `json:"project3"`
	Project4    string `json:"project4"`
	Project5    string `json:"project5"`
}
