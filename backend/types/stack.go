package types

import "database/sql"

type Stack struct {
	Base
	Name         string         `db:"name"`
	PathToFile   string         `db:"path_to_file"`
	Webhook      sql.NullString `db:"webhook"`
	UserID       int            `db:"user_id"`
	RepositoryID sql.NullInt32  `db:"repository_id"`
	Branch       sql.NullString `db:"branch"`
	Roles        sql.NullString `db:"roles"`
}

type StackData struct {
	Name       string `json:"name" validate:"required"`
	Webhook    string `json:"webhook"`
	Repo       string `json:"repo" validate:"required_without=Content"`
	RepoToken  string `json:"repo_token"`
	FileInRepo string `json:"file_in_repo"`
	Content    string `json:"content" validate:"required_without=Repo"`
}
