package handlers

import (
	"database/sql"
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
		LastName: sql.NullString{
			Valid:  true,
			String: data.LastName,
		},
		Role: data.Role,
		Email: sql.NullString{
			Valid:  true,
			String: data.Email,
		},
	}

	result, err := db.DB.Exec("INSERT INTO users (username, password_hash, first_name, last_name, role, email) VALUES (?, ?, ?, ?, ?, ?)", user.Username, user.PasswordHash, user.FirstName, user.LastName, user.Role, user.Email)
	if err != nil {
		log.Println(err)
		return c.String(500, "Server error")
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.String(500, "Server error")
	}

	return c.Redirect(302, "/users")
}

func (h UserHandler) CreateUserForm(c echo.Context) error {
	return Render(c, 200, compusers.CreateUserForm())
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
			return Render(c, 500, components.ErrorPopup(err, false))
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
			return Render(c, 500, components.ErrorPopup(err, false))
		}
	} else {
		sizeOfPageNum = 10
	}

	// Getting data
	role := c.(CustomContext).Locals["role"].(string)
	users := []models.User{}
	err := db.DB.Select(&users, "SELECT * FROM users LIMIT 10")
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return Render(c, 500, components.ErrorPopup(err, false))
	}
	var count int64
	err = db.DB.Get(&count, "SELECT count(*) FROM users")
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return Render(c, 500, components.ErrorPopup(err, false))
	}

	// Rendering response
	if c.Request().Header.Get("HX-Request") != "true" {
		return Render(c, 200, compusers.PageFull(users, pageNum, sizeOfPageNum, int(count), role))
	} else {
		return Render(c, 200, compusers.Page(users, pageNum, sizeOfPageNum, int(count), role))
	}
}
