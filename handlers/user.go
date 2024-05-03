package handlers

import (
	"log"
	"strconv"

	"github.com/RadeJR/containerama/components"
	compusers "github.com/RadeJR/containerama/components/users"
	"github.com/RadeJR/containerama/db"
	"github.com/RadeJR/containerama/models"
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
	return render(c, compusers.CreateUserForm())
}

func (h UserHandler) ShowUsers(c echo.Context) error {
	// PARSING QueryParam
	pageString := c.QueryParam("page")
	var pageNum int
	if pageString != "" {
		var err error
		pageNum, err = strconv.Atoi(pageString)
		if err != nil {
			c.Response().Header().Set("HX-Retarget", "#popup")
			return render(c, components.ErrorPopup(err, false))
		}
	} else {
		pageNum = 1
	}
	sizeOfPageString := c.QueryParam("size")
	var sizeOfPageNum int
	if sizeOfPageString != "" {
		var err error
		sizeOfPageNum, err = strconv.Atoi(sizeOfPageString)
		if err != nil {
			c.Response().Header().Set("HX-Retarget", "#popup")
			return render(c, components.ErrorPopup(err, false))
		}
	} else {
		sizeOfPageNum = 10
	}

	// Getting data
	role := c.(CustomContext).Locals["role"].(string)
	users := []models.User{}
	db.DB.Limit(10).Find(&users)
	var count int64
	db.DB.Find(&users).Count(&count)

	// Rendering response
	if c.Request().Header.Get("HX-Request") != "true" {
		return render(c, compusers.PageFull(users, pageNum, sizeOfPageNum, int(count), role))
	} else {
		render(c, components.Navbar(role, "Users"))
		return render(c, compusers.Page(users, pageNum, sizeOfPageNum, int(count)))
	}
}
