package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/RadeJR/containerama/components"
	"github.com/RadeJR/containerama/db"
	"github.com/RadeJR/containerama/models"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) error {
	res, err := db.DB.Exec("INSERT INTO users (username, email, first_name, last_name, password_hash, role) values (?, ?, ?, ?, ?, ?)", user.Username, user.Email, user.FirstName, user.LastName, user.PasswordHash, user.Role)
	if err != nil {
		slog.Error("Query failed")
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		slog.Error("Error extracting affected rows")
		return err
	}
	if rowsAffected < 1 {
		slog.Error("Rows affected by query below 1", "rowsAffected", string(rune(rowsAffected)))
		return errors.New("User was not created")
	}
	return nil
}

func EnsureAdminUserExists() {
	var count int
	db.DB.Get(&count, "SELECT count(*) FROM users WHERE role = ?", "admin")
	if count < 1 {
		slog.Info("Admin user wasnt found in a database, creating default one..")
		password, err := password.Generate(24, 8, 8, false, false)
		if err != nil {
			panic(err)
		}
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			panic(err)
		}
		user := models.User{
			Username:  "admin",
			FirstName: "Admin",
			LastName: sql.NullString{
				Valid:  true,
				String: "Admin",
			},
			Email:        sql.NullString{Valid: false},
			Role:         "admin",
			PasswordHash: string(passwordHash),
		}
		err = CreateUser(user)
		if err != nil {
			panic(err)
		}
		fmt.Println(password)
		slog.Info("Admin user created", "Username", user.Username, "Password", password)
	}
}

func GetUserByUsername(username string) (*models.User, error) {
	var user *models.User = &models.User{}
	err := db.DB.Get(user, "SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		slog.Error("Error retrieving user", "error", err)
		return nil, err
	}
	return user, nil
}

func NewUserRowData(user models.User) components.RowData {
	rowData := components.RowData{
		Fields: make([]string, 6),
	}

	rowData.Fields[0] = fmt.Sprint(user.ID)
	rowData.Fields[1] = user.Username
	if user.LastName.Valid {
		rowData.Fields[2] = fmt.Sprintf("%v %v", user.FirstName, user.LastName.String)
	} else {
		rowData.Fields[2] = user.FirstName
	}
	if user.Email.Valid {
		rowData.Fields[3] = user.Email.String
	} else {
		rowData.Fields[3] = "N/A"
	}
	rowData.Fields[4] = user.Role
	rowData.Fields[5] = user.CreatedAt.Format(time.ANSIC)
	return rowData
}
