package app

import (
	"github.com/MasLazu/CheatChatV2/controller"
	"github.com/MasLazu/CheatChatV2/middleware"
	"github.com/MasLazu/CheatChatV2/websocketProvider"

	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(
	router *mux.Router,
	corsMiddleware *middleware.CorsMiddleware,
	guestOnlyMiddlware *middleware.GuestOnlyMiddleware,
	loginOnlyMiddlware *middleware.LoginOnlyMiddleware,
	panicRecoveryMiddlware *middleware.PanicRecoveryMiddleware,
	userController *controller.UserController,
	groupController *controller.GroupController,
	contactController *controller.ContactController,
	chatController *controller.ChatController,
	websocketManager *websocketProvider.Manager,
) http.Handler {
	apiRoute := router.PathPrefix("/api").Subrouter()

	//panic recovery middleware
	apiRoute.Use(panicRecoveryMiddlware.MiddlewareFunc)

	//login only route
	loginRoute := apiRoute.PathPrefix("/login").Subrouter()
	loginRoute.Use(loginOnlyMiddlware.MiddlewareFunc)

	loginRoute.HandleFunc("/current", userController.Current).Methods("GET")

	loginRoute.HandleFunc("/group", groupController.Make).Methods("POST")
	loginRoute.HandleFunc("/groups", groupController.GetUserGroups).Methods("GET")

	loginRoute.HandleFunc("/contact", contactController.Add).Methods("POST")
	loginRoute.HandleFunc("/contacts", contactController.GetUserContacts).Methods("GET")

	loginRoute.HandleFunc("/chats/preview", chatController.GetPreviews).Methods("GET")
	loginRoute.HandleFunc("/chats/personal/{email}", chatController.GetPersonals).Methods("GET")
	loginRoute.HandleFunc("/chats/group/{id:[0-9]+}", chatController.GetGroups).Methods("GET")

	loginRoute.HandleFunc("/ws", websocketManager.Connect).Methods("GET")

	//guest only route
	guestRoute := apiRoute.PathPrefix("/guest").Subrouter()
	guestRoute.Use(guestOnlyMiddlware.MiddlewareFunc)

	guestRoute.HandleFunc("/register", userController.Register).Methods("POST")
	guestRoute.HandleFunc("/login", userController.Login).Methods("POST")

	//cors middleware
	return corsMiddleware.MiddlewareFunc(router)
}
