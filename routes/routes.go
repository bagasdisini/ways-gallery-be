package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	PostRoutes(r)
	AuthRoutes(r)
	TransactionRoutes(r)
}
