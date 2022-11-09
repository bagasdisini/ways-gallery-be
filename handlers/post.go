package handlers

import (
	postdto "backend/dto/post"
	dto "backend/dto/result"
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerPost struct {
	PostRepository repositories.PostRepository
}

func HandlerPost(PostRepository repositories.PostRepository) *handlerPost {
	return &handlerPost{PostRepository}
}

func (h *handlerPost) ShowPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	post, err := h.PostRepository.ShowPosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Post1
		post[i].Post1 = imagePath
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Post2
		post[i].Post2 = imagePath
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Post3
		post[i].Post3 = imagePath
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Post4
		post[i].Post4 = imagePath
	}

	for i, p := range post {
		imagePath := os.Getenv("PATH_FILE") + p.Post5
		post[i].Post5 = imagePath
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: http.StatusOK, Data: post}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPost) GetPostByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var post models.Post
	post, err := h.PostRepository.GetPostByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	post.Post1 = os.Getenv("PATH_FILE") + post.Post1
	post.Post2 = os.Getenv("PATH_FILE") + post.Post2
	post.Post3 = os.Getenv("PATH_FILE") + post.Post3
	post.Post4 = os.Getenv("PATH_FILE") + post.Post4
	post.Post5 = os.Getenv("PATH_FILE") + post.Post5

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: http.StatusOK, Data: post}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPost) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	adminInfo := r.Context().Value("authInfo").(jwt.MapClaims)
	userid := int(adminInfo["id"].(float64))
	dataUpload := r.Context().Value("dataFile")
	filepath := ""

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	if dataUpload != nil {
		filepath = dataUpload.(string)
	}

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err2 := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "waysgallery"})

	if err2 != nil {
		fmt.Println(err2.Error())
	}

	input := time.Now()

	dateParse := input.Format("2 Jan 2006 15:04")

	request := postdto.CreatePostRequest{
		Title: r.FormValue("title"),
		Desc:  r.FormValue("desc"),
		Date:  dateParse,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	post := models.Post{
		UserID: userid,
		Title:  request.Title,
		Desc:   request.Desc,
		Date:   request.Date,
		Post1:  resp.SecureURL,
	}

	post, err = h.PostRepository.CreatePost(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	post, err = h.PostRepository.GetPostByID(post.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: http.StatusOK, Data: post}
	json.NewEncoder(w).Encode(response)
}
