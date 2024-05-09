package services

import (
	"github.com/RadeJR/containerama/db"
	"github.com/RadeJR/containerama/models"
)

func GetStacks(user_id int, role string, page int, size int) ([]models.Stack, error) {
	stacks := make([]models.Stack, size)
	var err error
	offset := (page - 1) * size

	if role == "admin" {
		err = db.DB.Select(stacks, "SELECT * FROM stacks LIMIT ? OFFSET ?", size, offset)
	} else {
		err = db.DB.Select(stacks, "SELECT * FROM stacks WHERE user_id = ? LIMIT ? OFFSET ?", user_id, size, offset)
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
