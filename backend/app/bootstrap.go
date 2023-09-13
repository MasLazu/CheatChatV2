//go:build wireinject
// +build wireinject

package app

import (
	"net/http"

	"github.com/MasLazu/CheatChatV2/controller"
	"github.com/MasLazu/CheatChatV2/database"
	"github.com/MasLazu/CheatChatV2/middleware"
	"github.com/MasLazu/CheatChatV2/repository"
	"github.com/MasLazu/CheatChatV2/service"
	"github.com/MasLazu/CheatChatV2/websocketProvider"
	"github.com/google/wire"
	"github.com/gorilla/mux"
)

func BootstrapApp() http.Handler {
	wire.Build(
		//get database connection
		database.GetDBConn,

		//repository
		repository.NewChatsRepository,
		repository.NewContactRepository,
		repository.NewGroupRepository,
		repository.NewPersonalRepository,
		repository.NewSessionRepository,
		repository.NewUserRepository,

		//service
		service.NewChatService,
		service.NewContactService,
		service.NewGroupService,
		service.NewSessionService,
		service.NewUserService,

		//middleware
		middleware.NewCorsMiddleware,
		middleware.NewGuestOnlyMiddleware,
		middleware.NewLoginOnlyMiddleware,
		middleware.NewPanicRecoveryMiddleware,

		//websocket
		websocketProvider.NewManager,

		//controller
		controller.NewChatController,
		controller.NewContactController,
		controller.NewGroupController,
		controller.NewUserController,

		//router
		mux.NewRouter,
		NewRouter,
	)

	return nil
}
