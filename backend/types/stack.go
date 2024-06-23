package types

type Stack struct {
	Base
	Name         string `db:"name"`
	PathToFile   string `db:"path_to_file"`
	Webhook      string `db:"webhook"`
	UserID       int    `db:"user_id"`
	RepositoryID int    `db:"repository_id"`
	Branch       string `db:"branch"`
}
