package types

type Repository struct {
	Base
	Name           string `db:"name"`
	EncryptedToken string `db:"encrypted_token"`
	URL            string `db:"url"`
	UserID         int    `db:"user_id"`
}
