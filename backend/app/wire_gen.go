// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/MasLazu/CheatChatV2/controller"
	"github.com/MasLazu/CheatChatV2/database"
	"github.com/MasLazu/CheatChatV2/middleware"
	"github.com/MasLazu/CheatChatV2/repository"
	"github.com/MasLazu/CheatChatV2/service"
	"github.com/MasLazu/CheatChatV2/websocketProvider"
	"github.com/gorilla/mux"
	"net/http"
)

// Injectors from bootstrap.go:

func BootstrapApp() http.Handler {
	router := mux.NewRouter()
	corsMiddleware := middleware.NewCorsMiddleware()
	db := database.GetDBConn()
	userRepository := repository.NewUserRepository(db)
	sessionRepository := repository.NewSessionRepository(db)
	sessionService := service.NewSessionService(userRepository, sessionRepository)
	guestOnlyMiddleware := middleware.NewGuestOnlyMiddleware(sessionService)
	loginOnlyMiddleware := middleware.NewLoginOnlyMiddleware(sessionService)
	panicRecoveryMiddleware := middleware.NewPanicRecoveryMiddleware()
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(sessionService, userService)
	groupRepository := repository.NewGroupRepository(db)
	groupService := service.NewGroupService(groupRepository)
	groupController := controller.NewGroupController(sessionService, groupService, groupRepository)
	personalRepository := repository.NewPersonalRepository(db)
	contactRepository := repository.NewContactRepository(db)
	contactService := service.NewContactService(userRepository, personalRepository, contactRepository)
	contactController := controller.NewContactController(sessionService, contactService, contactRepository)
	chatRepository := repository.NewChatsRepository(db)
	chatController := controller.NewChatController(sessionService, chatRepository)
	chatService := service.NewChatService(chatRepository, personalRepository, groupRepository)
	manager := websocketProvider.NewManager(sessionService, chatService, groupRepository)
	handler := NewRouter(router, corsMiddleware, guestOnlyMiddleware, loginOnlyMiddleware, panicRecoveryMiddleware, userController, groupController, contactController, chatController, manager)
	return handler
}
