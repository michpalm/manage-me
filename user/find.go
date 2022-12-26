package user

import (
	"github.com/michpalm/manage-me/database"
	"github.com/michpalm/manage-me/database/models"
)

type EmailNotExistsError struct{}

func (*EmailNotExistsError) Error() string {
	return "email not exists"
}

func FindByEmail(email string) (*models.User, error) {
	var user models.User
	res := database.DB.Db.Find(&user, &models.User{Email: email})
	return &user, res.Error
}
