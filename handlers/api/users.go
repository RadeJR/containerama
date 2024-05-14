package api

import (
	"database/sql"
	"net/http"

	"github.com/RadeJR/containerama/db"
	"github.com/RadeJR/containerama/models"
	"github.com/RadeJR/containerama/services"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c echo.Context) error {
	type userData struct {
		Username  string `form:"username" validate:"required,alphanum"`
		Password  string `form:"password" validate:"required"`
		FirstName string `form:"firstname" validate:"required"`
		LastName  string `form:"lastname"`
		Role      string `form:"role"`
		Email     string `form:"email" validate:"email"`
	}
	data := userData{}
	if err := c.Bind(&data); err != nil {
		return err
	}

	if err := services.Validate.Struct(data); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return err
	}

	user := models.User{
		Username:     data.Username,
		PasswordHash: string(hashedPassword),
		FirstName:    data.FirstName,
		Role:         data.Role,
	}
	if data.LastName != "" {
		user.LastName = sql.NullString{
			Valid:  true,
			String: data.LastName,
		}
	}
	if data.Email != "" {
		user.Email = sql.NullString{
			Valid:  true,
			String: data.Email,
		}
	}

	result, err := db.DB.Exec("INSERT INTO users (username, password_hash, first_name, last_name, role, email) VALUES (?, ?, ?, ?, ?, ?)", user.Username, user.PasswordHash, user.FirstName, user.LastName, user.Role, user.Email)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "User wasn't created")
	}

	return c.JSON(http.StatusCreated, "User created")
}

func ShowUsers(c echo.Context) error {
	users := []models.User{}
	err := db.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}
