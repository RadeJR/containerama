package handler

import (
	"log"

	"github.com/RadeJR/itcontainers/model"
	"github.com/RadeJR/itcontainers/view/user"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

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

	user := model.User{
		Username:     data.Username,
		PasswordHash: string(hashedPassword),
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		Role:         data.Role,
		Email:        data.Email,
	}

	if err := h.DB.Create(&user).Error; err != nil {
		log.Println(err)
		return c.String(500, "Server error")
	}

	return c.Redirect(302, "/users")
}

func (h UserHandler) CreateUserForm(c echo.Context) error {
	return render(c, user.CreateUserForm())
}

func (h UserHandler) ShowUsers(c echo.Context) error {
	users := []model.User{}
	h.DB.Limit(10).Find(&users)
	var count int64
	h.DB.Find(&users).Count(&count)
	log.Println(count)
	return render(c, user.UserPage(users, count))
}
