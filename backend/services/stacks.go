package services

import (
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/RadeJR/containerama/db"
	"github.com/RadeJR/containerama/types"
)

func GetStacks(userID string, roles []string) ([]types.Stack, error) {
	var result []types.Stack
	var err error
	for _, v := range roles {
		if v == "admin" {
			err = db.DB.Select(&result, "SELECT * FROM stacks")
			if err != nil {
				return nil, err
			}
			return result, nil
		}
	}
	var stacks []types.Stack

	// get stack that user owns
	err = db.DB.Select(&result, "SELECT * FROM stacks WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	// get stack that are owned by roles
	for _, v := range roles {
		err = db.DB.Select(&stacks, "SELECT * FROM stacks WHERE roles LIKE '%?%'", v)
		if err != nil {
			return nil, err
		}
		result = append(result, stacks...)
	}
	return result, nil
}

func CreateStack(data types.StackData, userID string) error {
	path := "data/stacks/" + data.Name + "/docker-compose.yml"
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		slog.Error("Error creating directories")
		return err
	}
	err = os.WriteFile(path, []byte(data.Content), 0644)
	if err != nil {
		slog.Error("Error creating file", "error", err)
		return err
	}

	pathDir := filepath.Dir(path)
	args := "compose --project-directory " + pathDir + " up -d"
	cmd := exec.Command("docker", strings.Split(args, " ")...)
	_, err = cmd.Output()
	if err != nil {
		return err
	}

	_, err = db.DB.Exec("INSERT INTO stacks (name, path_to_file, user_id, webhook ) VALUES(?, ?, ?, ?)", data.Name, path, userID, data.Webhook)
	if err != nil {
		return err
	}

	return nil
}

func DeleteStack(name string, userID string) error {
	stack := types.Stack{}
	err := db.DB.Get(&stack, "SELECT * FROM stacks WHERE user_id = ? AND name = ?", userID, name)
	if err != nil {
		return err
	}

	pathDir := filepath.Dir(stack.PathToFile)
	args := "compose --project-directory " + pathDir + " down"
	cmd := exec.Command("docker", strings.Split(args, " ")...)
	_, err = cmd.Output()
	if err != nil {
		return err
	}

	err = os.RemoveAll(pathDir)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec("DELETE FROM stacks WHERE name = ?", name)
	if err != nil {
		return err
	}

	return nil
}
