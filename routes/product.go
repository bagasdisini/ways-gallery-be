package routes

import (
	"backend/handlers"
	"backend/pkg/middleware"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func PostRoutes(r *mux.Router) {
	postRepository := repositories.RepositoryPost(mysql.DB)
	h := handlers.HandlerPost(postRepository)

	r.HandleFunc("/posts", middleware.Auth(h.ShowPosts)).Methods("GET")
	r.HandleFunc("/post/{id}", middleware.Auth(h.GetPostByID)).Methods("GET")
	r.HandleFunc("/post", middleware.Auth(middleware.UploadImage(h.CreatePost))).Methods("POST")
}
