package handlers

import (
	postdto "backend/dto/post"
	dto "backend/dto/result"
	"backend/models"
	"backend/repositories"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

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

	Image, err := h.PostRepository.ShowPosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range Image {
		imagePath := os.Getenv("PATH_FILE") + p.Image1
		Image[i].Image1 = imagePath
	}

	for i, p := range Image {
		imagePath := os.Getenv("PATH_FILE") + p.Image2
		Image[i].Image2 = imagePath
	}

	for i, p := range Image {
		imagePath := os.Getenv("PATH_FILE") + p.Image3
		Image[i].Image3 = imagePath
	}

	for i, p := range Image {
		imagePath := os.Getenv("PATH_FILE") + p.Image4
		Image[i].Image4 = imagePath
	}

	for i, p := range Image {
		imagePath := os.Getenv("PATH_FILE") + p.Image5
		Image[i].Image5 = imagePath
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: http.StatusOK, Data: Image}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPost) GetPostByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var Image models.Post
	Image, err := h.PostRepository.GetPostByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	Image.Image1 = os.Getenv("PATH_FILE") + Image.Image1
	Image.Image2 = os.Getenv("PATH_FILE") + Image.Image2
	Image.Image3 = os.Getenv("PATH_FILE") + Image.Image3
	Image.Image4 = os.Getenv("PATH_FILE") + Image.Image4
	Image.Image5 = os.Getenv("PATH_FILE") + Image.Image5

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: http.StatusOK, Data: Image}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerPost) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	adminInfo := r.Context().Value("authInfo").(jwt.MapClaims)
	userid := int(adminInfo["id"].(float64))
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
		Image1: resp.SecureURL,
		Image2: resp2.SecureURL,
		Image3: resp3.SecureURL,
		Image4: resp4.SecureURL,
		Image5: resp5.SecureURL,
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
