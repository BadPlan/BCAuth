package repositories

import (
	"BCAuth/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SessionRepository struct {
	tx *gorm.DB
}

func (s SessionRepository) CreateSession(ctx *gin.Context, session models.Session) (models.Session, error) {
	if err := s.tx.Omit("Users").Create(&session).Error; err != nil {
		return models.Session{}, err
	}
	return session, nil
}

func (s SessionRepository) UserByToken(ctx *gin.Context, session models.Session) (models.Session, error) {
	if err := s.tx.Where(&session).First(&session).Error; err != nil {
		return models.Session{}, err
	}
	var user models.User
	if err := s.tx.Model(&models.User{}).Where("id = ?", session.UserID).First(&user).Error; err != nil {
		return models.Session{}, err
	}
	session.User = user
	if err := s.tx.Model(&user).Association("Roles").Find(&session.User.Roles); err != nil {
		return models.Session{}, err
	}
	return session, nil
}

func InitSessionRepository(database *gorm.DB) *SessionRepository {
	return &SessionRepository{
		tx: database.Session(&gorm.Session{SkipDefaultTransaction: true}),
	}
}
