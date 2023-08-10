package router

import (
	"github.com/MasLazu/CheatChatV2/controller"
	"github.com/MasLazu/CheatChatV2/middleware"
	"github.com/MasLazu/CheatChatV2/websocketProvider"

	"github.com/gorilla/mux"
)

func MainRouter(router *mux.Router) {
	websocketManager := websocketProvider.NewWebsocketManager()

	loginRoute := router.PathPrefix("/login").Subrouter()
	loginRoute.Use(middleware.LoginOnlyMiddleware)

	loginRoute.HandleFunc("/current", controller.CurrentController).Methods("GET")

	loginRoute.HandleFunc("/group", controller.MakeGroupController).Methods("POST")
	loginRoute.HandleFunc("/groups", controller.GetUserGroupsController).Methods("GET")

	loginRoute.HandleFunc("/contact", controller.AddContactController).Methods("POST")
	loginRoute.HandleFunc("/contacts", controller.GetContactUserController).Methods("GET")

	loginRoute.HandleFunc("/chats/preview", controller.GetPreviewChatController).Methods("GET")
	loginRoute.HandleFunc("/chats/personal/{email}", controller.GetPersonalFullChat).Methods("GET")
	loginRoute.HandleFunc("/chats/group/{id:[0-9]+}", controller.GetGroupFullChats).Methods("GET")

	loginRoute.HandleFunc("/ws", websocketManager.Connect).Methods("GET")

	guestRoute := router.PathPrefix("/guest").Subrouter()
	guestRoute.Use(middleware.GuestOnlyMiddleware)

	guestRoute.HandleFunc("/register", controller.RegisterController).Methods("POST")

	guestRoute.HandleFunc("/login", controller.LoginController).Methods("POST")
}
