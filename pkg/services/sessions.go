package services

import (
	"BCAuth/internal/models"
	"BCAuth/pkg/repositories"
	"github.com/gin-gonic/gin"
)

type SessionService struct {
	sessionRepo repositories.Session
}

func (s SessionService) CreateSession(ctx *gin.Context, session models.Session) (models.Session, error) {
	session, err := s.sessionRepo.CreateSession(ctx, session)
	if err != nil {
		return models.Session{}, err
	}
	return session, nil
}

func (s SessionService) UserByToken(ctx *gin.Context, token string) (models.UserInfo, error) {
	var session models.Session
	session.Token = new(string)
	*session.Token = token
	found, err := s.sessionRepo.UserByToken(ctx, session)
	if err != nil {
		return models.UserInfo{}, err
	}
	var user models.UserInfo
	user.ID = &found.User.ID
	user.Name = found.User.Name
	user.Email = found.User.Email
	var roles []models.RoleInfo
	for _, r := range found.User.Roles {
		role := models.RoleInfo{}
		role.ID = new(uint)
		*role.ID = r.ID
		role.Name = r.Name
		role.Description = r.Description
		roles = append(roles, role)
	}
	user.Roles = roles
	return user, err
}

func InitSessionService(repository *repositories.Session) *SessionService {
	return &SessionService{
		sessionRepo: *repository,
	}
}
