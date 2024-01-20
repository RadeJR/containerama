package handlers

import (
	"log"

	"github.com/RadeJR/itcontainers/components"
	"github.com/RadeJR/itcontainers/db"
	"github.com/RadeJR/itcontainers/models"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct{}

func (h UserHandler) CreateUser(c echo.Context) error {
	type userData struct {
		Username  string `form:"username"`
		Password  string `form:"password"`
		FirstName string `form:"firstname"`
		LastName  string `form:"lastname"`
		Role      string `form:"role"`
		Email     string `form:"email"`
	}
	data := userData{}
	c.Bind(&data)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return c.String(500, "Server error")
	}

	user := models.User{
		Username:     data.Username,
		PasswordHash: string(hashedPassword),
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		Role:         data.Role,
		Email:        data.Email,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		log.Println(err)
		return c.String(500, "Server error")
	}

	return c.Redirect(302, "/users")
}

func (h UserHandler) CreateUserForm(c echo.Context) error {
	return render(c, components.CreateUserForm())
}

func (h UserHandler) ShowUsers(c echo.Context) error {
	users := []models.User{}
	db.DB.Limit(10).Find(&users)
	var count int64
	db.DB.Find(&users).Count(&count)
	log.Println(count)
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(500, "Server error")
	}
	return render(c, components.UserPage(users, count, sess.Values["role"].(string)))
}
