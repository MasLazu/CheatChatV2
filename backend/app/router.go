package app

import (
	"github.com/MasLazu/CheatChatV2/controller"
	"github.com/MasLazu/CheatChatV2/middleware"
	"github.com/MasLazu/CheatChatV2/websocketProvider"

	"net/http"

	"github.com/gorilla/mux"
)

type Router interface {
	Handler() http.Handler
}

type RouterImpl struct {
	router                 *mux.Router
	corsMiddleware         *middleware.CorsMiddleware
	guestOnlyMiddlware     *middleware.GuestOnlyMiddleware
	loginOnlyMiddlware     *middleware.LoginOnlyMiddleware
	panicRecoveryMiddlware *middleware.PanicRecoveryMiddleware
	userController         *controller.UserController
	groupController        *controller.GroupController
	contactController      *controller.ContactController
	chatController         *controller.ChatController
	websocketManager       *websocketProvider.Manager
}

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
) Router {
	return &RouterImpl{
		router:                 router,
		corsMiddleware:         corsMiddleware,
		guestOnlyMiddlware:     guestOnlyMiddlware,
		loginOnlyMiddlware:     loginOnlyMiddlware,
		panicRecoveryMiddlware: panicRecoveryMiddlware,
		userController:         userController,
		groupController:        groupController,
		contactController:      contactController,
		chatController:         chatController,
		websocketManager:       websocketManager,
	}
}

func (router *RouterImpl) Handler() http.Handler {
	apiRoute := router.router.PathPrefix("/api").Subrouter()

	//panic recovery middleware
	apiRoute.Use(router.panicRecoveryMiddlware.MiddlewareFunc)

	//login only route
	loginRoute := apiRoute.PathPrefix("/login").Subrouter()
	loginRoute.Use(router.loginOnlyMiddlware.MiddlewareFunc)

	loginRoute.HandleFunc("/current", router.userController.Current).Methods("GET")

	loginRoute.HandleFunc("/group", router.groupController.Make).Methods("POST")
	loginRoute.HandleFunc("/groups", router.groupController.GetUserGroups).Methods("GET")

	loginRoute.HandleFunc("/contact", router.contactController.Add).Methods("POST")
	loginRoute.HandleFunc("/contacts", router.contactController.GetUserContacts).Methods("GET")

	loginRoute.HandleFunc("/chats/preview", router.chatController.GetPreviews).Methods("GET")
	loginRoute.HandleFunc("/chats/personal/{email}", router.chatController.GetPersonals).Methods("GET")
	loginRoute.HandleFunc("/chats/group/{id:[0-9]+}", router.chatController.GetGroups).Methods("GET")

	loginRoute.HandleFunc("/ws", router.websocketManager.Connect).Methods("GET")

	//guest only route
	guestRoute := apiRoute.PathPrefix("/guest").Subrouter()
	guestRoute.Use(router.guestOnlyMiddlware.MiddlewareFunc)

	guestRoute.HandleFunc("/register", router.userController.Register).Methods("POST")
	guestRoute.HandleFunc("/login", router.userController.Login).Methods("POST")

	//cors middleware
	return router.corsMiddleware.MiddlewareFunc(router.router)
}
