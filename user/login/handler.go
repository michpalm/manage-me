package login

import (
	"github.com/gofiber/fiber/v2"
	"github.com/michpalm/manage-me/database"
	"github.com/michpalm/manage-me/database/models"
	"github.com/michpalm/manage-me/user"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Request struct {
	Email    string
	Password string
}

type Response struct {
	User *models.User
}

type PasswordMismatchError struct{}

func (e *PasswordMismatchError) Error() string {
	return "password didn't match"
}

func NewView(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title":    "Welcome to this app",
		"Subtitle": "Use your credentials to log in!",
	})
}

func Login(c *fiber.Ctx) error {
	request := new(Request)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	existingUser, err := user.FindByEmail(request.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(request.Password))
	if err != nil {
		return &PasswordMismatchError{}
	}

	session := &models.Session{
		Username:  existingUser.Username,
		Timestamp: time.Now(),
	}

	err = database.DB.Db.Create(&session).Error
	if err != nil {
		return err
	}

	sessions, err := user.GetLatestSessions(existingUser.Username)
	if err != nil {
		return err
	}

	return c.Render("home", fiber.Map{
		"Title":    "Welcome " + existingUser.Username + "!",
		"Subtitle": "Here are your 10 latest sessions:",
		"Sessions": sessions,
	})
}
