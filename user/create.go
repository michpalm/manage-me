package user

import (
	"github.com/michpalm/manage-me/database"
	"github.com/michpalm/manage-me/database/models"
)

func Create(user *models.User) error {
	err := database.DB.Db.Create(user).Error
	if err != nil {
		if models.IsUniqueConstraintError(err, models.UniqueConstraintUsername) {
			return &models.UsernameDuplicateError{Username: user.Username}
		}
		if models.IsUniqueConstraintError(err, models.UniqueConstraintEmail) {
			return &models.EmailDuplicateError{Email: user.Email}
		}
		return err
	}
	return nil
}
