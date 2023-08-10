package service

import (
	"context"
	"crypto/rand"
	"errors"
	"github.com/MasLazu/CheatChatV2/model"
	"github.com/MasLazu/CheatChatV2/model/domain"
	"github.com/MasLazu/CheatChatV2/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type SessionService interface {
	Login(request model.LoginUserRequest, ctx context.Context) (domain.Session, error)
	Current(request *http.Request, ctx context.Context) (domain.User, error)
}

type SessionServiceImpl struct {
}

func NewSessionService() SessionService {
	return &SessionServiceImpl{}
}

func (service SessionServiceImpl) Login(request model.LoginUserRequest, ctx context.Context) (domain.Session, error) {
	var session domain.Session

	userRepository := repository.NewUsersRepository()
	user, err := userRepository.GetByEmail(ctx, request.Email)
	if err != nil {
		return session, errors.New("credential not metch")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return session, errors.New("credential not metch")
	}

	byteToken := make([]byte, 32)
	if _, err := rand.Read(byteToken); err != nil {
		return session, errors.New("something went wrong")
	}

	session.UserEmail = user.Email
	session.Token = uuid.New().String()
	session.ExpiredAt = time.Now().Add(24 * time.Hour)

	sessionRepository := repository.NewSessionRepository()
	if err := sessionRepository.InsertIfExistUpdate(session, ctx); err != nil {
		return session, errors.New("something went wrong")
	}

	return session, nil
}

func (service SessionServiceImpl) Current(request *http.Request, ctx context.Context) (domain.User, error) {
	user := domain.User{}
	sessionCookie, err := request.Cookie("session")
	if err != nil {
		return user, err
	}

	sessionRepository := repository.NewSessionRepository()
	currentSession, err := sessionRepository.GetByToken(ctx, sessionCookie.Value)
	if err != nil {
		return user, err
	}

	if currentSession.ExpiredAt.Before(time.Now()) {
		return user, err
	}

	userRepository := repository.NewUsersRepository()
	user, err = userRepository.GetByEmail(ctx, currentSession.UserEmail)
	if err != nil {
		return user, err
	}

	return user, nil
}
