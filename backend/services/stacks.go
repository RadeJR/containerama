package services

import (
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/RadeJR/containerama/db"
	"github.com/RadeJR/containerama/types"
)

func GetStacks(userID string, isAdmin bool) ([]types.Stack, error) {
	var stacks []types.Stack
	var err error

	if isAdmin {
		err = db.DB.Select(&stacks, "SELECT * FROM stacks")
	} else {
		err = db.DB.Select(&stacks, "SELECT * FROM stacks WHERE user_id = ?", userID)
	}
	if err != nil {
		return nil, err
	}
	return stacks, nil
}

func CreateStackFromFile(name string, userID string, filePath string) error {
	pathDir := filepath.Dir(filePath)
	args := "compose --project-directory " + pathDir + " up -d"
	cmd := exec.Command("docker", strings.Split(args, " ")...)
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	_, err = db.DB.Exec("INSERT INTO stacks (name, path_to_file, user_id ) VALUES(?, ?, ?)", name, filePath, userID)
	if err != nil {
		return err
	}

	return nil
}
