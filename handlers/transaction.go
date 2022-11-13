package handlers

import (
	dto "backend/dto/result"
	transactiondto "backend/dto/transaction"
	"backend/models"
	"backend/repositories"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
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
	p := r.FormValue("admin_id")
	adminP, _ := strconv.Atoi(p)

	request := transactiondto.CreateTransactionRequest{
		AdminID:   adminP,
		Desc:      r.FormValue("desc"),
		Title:     r.FormValue("title"),
		StartDate: r.FormValue("startDate"),
		EndDate:   r.FormValue("endDate"),
		Price:     r.FormValue("price"),
		Status:    "pending",
	}

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
		Status:    request.Status,
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

	dataUpload := r.Context().Value("dataPost")
	dataUpload2 := r.Context().Value("dataPost2")
	dataUpload3 := r.Context().Value("dataPost3")
	dataUpload4 := r.Context().Value("dataPost4")
	dataUpload5 := r.Context().Value("dataPost5")
	filepath := ""
	filepath2 := ""
	filepath3 := ""
	filepath4 := ""
	filepath5 := ""

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	if dataUpload != nil {
		filepath = dataUpload.(string)
	}
	if dataUpload2 != nil {
		filepath2 = dataUpload2.(string)
	}
	if dataUpload3 != nil {
		filepath3 = dataUpload3.(string)
	}
	if dataUpload4 != nil {
		filepath4 = dataUpload4.(string)
	}
	if dataUpload5 != nil {
		filepath5 = dataUpload5.(string)
	}

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err2 := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "waysgallery"})
	resp2, _ := cld.Upload.Upload(ctx, filepath2, uploader.UploadParams{Folder: "waysgallery"})
	resp3, _ := cld.Upload.Upload(ctx, filepath3, uploader.UploadParams{Folder: "waysgallery"})
	resp4, _ := cld.Upload.Upload(ctx, filepath4, uploader.UploadParams{Folder: "waysgallery"})
	resp5, _ := cld.Upload.Upload(ctx, filepath5, uploader.UploadParams{Folder: "waysgallery"})

	if err2 != nil {
		fmt.Println(err2.Error())
	}

	request := transactiondto.UpdateTransactionRequest{
		Status:      r.FormValue("status"),
		ProjectDesc: r.FormValue("projectDesc"),
	}

	transaction := models.Transaction{}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if request.Status != "" {
		transaction.Status = request.Status
	}

	if request.ProjectDesc != "" {
		transaction.ProjectDesc = request.ProjectDesc
	}

	if filepath != "" {
		transaction.Image1 = resp.SecureURL
	}
	if filepath2 != "" {
		transaction.Image2 = resp2.SecureURL
	}
	if filepath3 != "" {
		transaction.Image3 = resp3.SecureURL
	}
	if filepath4 != "" {
		transaction.Image4 = resp4.SecureURL
	}
	if filepath5 != "" {
		transaction.Image5 = resp5.SecureURL
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
