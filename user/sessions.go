package user

import (
	"github.com/michpalm/manage-me/database"
	"github.com/michpalm/manage-me/database/models"
)

func GetLatestSessions(username string) (*[]models.Session, error) {
	var sessions []models.Session
	res := database.DB.Db.Order("timestamp desc").Limit(10).Where("username = ?", username).Find(&sessions)
	return &sessions, res.Error
}
