package handlers

import (
	dto "backend/dto/result"
	transactiondto "backend/dto/transaction"
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) ShowTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transaction, err := h.TransactionRepository.ShowTransaction()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range transaction {
		imagePath := os.Getenv("PATH_FILE") + p.Image1
		transaction[i].Image1 = imagePath
	}

	for i, p := range transaction {
		imagePath := os.Getenv("PATH_FILE") + p.Image2
		transaction[i].Image2 = imagePath
	}

	for i, p := range transaction {
		imagePath := os.Getenv("PATH_FILE") + p.Image3
		transaction[i].Image3 = imagePath
	}

	for i, p := range transaction {
		imagePath := os.Getenv("PATH_FILE") + p.Image4
		transaction[i].Image4 = imagePath
	}

	for i, p := range transaction {
		imagePath := os.Getenv("PATH_FILE") + p.Image5
		transaction[i].Image5 = imagePath
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var transaction models.Transaction
	transaction, err := h.TransactionRepository.GetTransactionByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction.Image1 = os.Getenv("PATH_FILE") + transaction.Image1
	transaction.Image2 = os.Getenv("PATH_FILE") + transaction.Image2
	transaction.Image3 = os.Getenv("PATH_FILE") + transaction.Image3
	transaction.Image4 = os.Getenv("PATH_FILE") + transaction.Image4
	transaction.Image5 = os.Getenv("PATH_FILE") + transaction.Image5

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	adminInfo := r.Context().Value("authInfo").(jwt.MapClaims)
	buyerId := int(adminInfo["id"].(float64))

	request := new(transactiondto.CreateTransactionRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction := models.Transaction{
		AdminID:   request.AdminID,
		BuyerID:   buyerId,
		Title:     request.Title,
		Desc:      request.Desc,
		StartDate: request.StartDate,
		EndDate:   request.EndDate,
		Price:     request.Price,
		Status:    "waiting",
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction, err = h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction, _ = h.TransactionRepository.GetTransactionByID(transaction.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: http.StatusOK, Data: transaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(transactiondto.UpdateTransactionRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction := models.Transaction{}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if request.Status != "" {
		transaction.Status = request.Status
	}

	if request.ProjectDesc != "" {
		transaction.ProjectDesc = request.ProjectDesc
	}

	if request.Image1 != "" {
		transaction.Image1 = request.Image1
	}

	if request.Image2 != "" {
		transaction.Image2 = request.Image2
	}

	if request.Image3 != "" {
		transaction.Image3 = request.Image3
	}

	if request.Image4 != "" {
		transaction.Image4 = request.Image4
	}

	if request.Image5 != "" {
		transaction.Image5 = request.Image5
	}

	data, err := h.TransactionRepository.UpdateTransaction(transaction, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	transaction, err := h.TransactionRepository.GetTransactionByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TransactionRepository.DeleteTransaction(transaction, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
