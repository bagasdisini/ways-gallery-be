package handlers

import (
	dto "backend/dto/result"
	usersdto "backend/dto/users"
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

func (h *handler) ShowUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := h.UserRepository.ShowUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	for i, p := range users {
		users[i].Image = "http://localhost:5000/uploads/" + p.Image
	}

	for i, p := range users {
		users[i].BestArt = "http://localhost:5000/uploads/" + p.BestArt
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUserByIDUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		response := dto.ErrorResult{Status: http.StatusNotFound, Message: "ID: " + strconv.Itoa(id) + " not found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	user.Image = "http://localhost:5000/uploads/" + user.Image
	user.BestArt = "http://localhost:5000/uploads/" + user.BestArt

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataImage := r.Context().Value("dataFile")
	dataBestArt := r.Context().Value("dataBestArt")
	filepath := ""
	fileArt := ""

	// var ctx = context.Background()
	// var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	// var API_KEY = os.Getenv("API_KEY")
	// var API_SECRET = os.Getenv("API_SECRET")

	if dataImage != nil {
		filepath = dataImage.(string)
	}

	if dataBestArt != nil {
		fileArt = dataBestArt.(string)
	}

	// cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// resp, err2 := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "waysgallery"})

	// if err2 != nil {
	// 	fmt.Println(err2.Error())
	// }

	request := usersdto.UpdateUserRequest{
		Name:      r.FormValue("name"),
		Greeting:  r.FormValue("greeting"),
		Following: r.FormValue("following"),
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user := models.User{}

	user.ID = id

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Greeting != "" {
		user.Greeting = request.Greeting
	}

	if request.Following != "" {
		user.Following = request.Following
	}

	if filepath != "" {
		user.Image = filepath
	}

	if fileArt != "" {
		user.BestArt = fileArt
	}

	data, err := h.UserRepository.UpdateUser(user, id)
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

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUserByIDUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.UserRepository.DeleteUser(user, id)
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
