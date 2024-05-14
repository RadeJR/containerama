package services

import (

	"github.com/RadeJR/containerama/db"
	"github.com/RadeJR/containerama/models"
)

func GetStacks(userID int, role string) ([]models.Stack, error) {
	var stacks []models.Stack
	var err error

	if role == "admin" {
		err = db.DB.Select(&stacks, "SELECT * FROM stacks")
	} else {
		err = db.DB.Select(&stacks, "SELECT * FROM stacks WHERE user_id = ?", userID)
	}
	if err != nil {
		return nil, err
	}
	return stacks, nil
}

func CountStacks(userId int, role string) (int, error) {
	var count int
	var err error

	if role == "admin" {
		err = db.DB.Get(count, "SELECT count(*) FROM stacks")
	} else {
		err = db.DB.Get(count, "SELECT count(*) FROM stacks WHERE user_id = ?", userId)
	}
	if err != nil {
		return 0, err
	}
	return count, nil
}
