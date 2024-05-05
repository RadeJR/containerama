package handlers

import (
	"database/sql"
	"net/http"

	"github.com/RadeJR/containerama/components"
	compusers "github.com/RadeJR/containerama/components/users"
	"github.com/RadeJR/containerama/db"
	"github.com/RadeJR/containerama/models"
	"github.com/RadeJR/containerama/services"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var Headers = []string{"ID", "Username", "Full Name", "Email", "Role", "Created at"}

type UserHandler struct{}

func (h UserHandler) CreateUser(c echo.Context) error {
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

	return c.Redirect(302, "/users")
}

func (h UserHandler) CreateUserForm(c echo.Context) error {
	return Render(c, 200, compusers.CreateUserForm())
}

func (h UserHandler) ShowUsers(c echo.Context) error {
	page, size, err := GetPaginationInfo(c)
	if err != nil {
		return err
	}
	// Getting data
	role := c.(CustomContext).Locals["role"].(string)
	users := []models.User{}
	err = db.DB.Select(&users, "SELECT * FROM users LIMIT 10")
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return Render(c, 500, components.ErrorPopup(err))
	}
	var count int64
	err = db.DB.Get(&count, "SELECT count(*) FROM users")
	if err != nil {
		c.Response().Header().Set("HX-Retarget", "#popup")
		return Render(c, 500, components.ErrorPopup(err))
	}

	tableData := components.TableData{
		Rows: make([]components.RowData, len(users)),
	}
	tableData.Headers = Headers
	for k,v := range users {
		tableData.Rows[k] = services.NewUserRowData(v)
	}

	// Rendering response
	if c.Request().Header.Get("HX-Request") != "true" {
		return Render(c, 200, compusers.PageFull(tableData, page, size, int(count), role))
	} else {
		return Render(c, 200, compusers.Page(tableData, page, size, int(count), role))
	}
}
