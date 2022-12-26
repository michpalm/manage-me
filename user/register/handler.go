package register

import (
	"github.com/gofiber/fiber/v2"
	"github.com/michpalm/manage-me/database"
	"github.com/michpalm/manage-me/database/models"
	"github.com/michpalm/manage-me/user"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Request struct {
	Username string
	Email    string
	Password string
}

type Response struct {
	Id uint
}

func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":    "Welcome to this app",
		"Subtitle": "Are you an existing user or a new one?",
		//"Facts":    facts, // send the facts to the view
	})
}

func NewView(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title":    "Welcome to this app",
		"Subtitle": "Fill in the form to register.",
		//"Facts":    facts, // send the facts to the view
	})
}

func Register(c *fiber.Ctx) error {
	request := new(Request)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser := &models.User{
		Username:     request.Username,
		Email:        request.Email,
		PasswordHash: string(passwordHash),
	}

	err = user.Create(newUser)
	if err != nil {
		return err
	}

	session := &models.Session{
		Username:  newUser.Username,
		Timestamp: time.Now(),
	}

	err = database.DB.Db.Create(&session).Error
	if err != nil {
		return err
	}

	sessions, err := user.GetLatestSessions(newUser.Username)
	if err != nil {
		return err
	}

	return c.Render("home", fiber.Map{
		"Title":    "Welcome " + newUser.Username + "!",
		"Subtitle": "Here are your 10 latest sessions:",
		"Sessions": sessions,
	})
}
