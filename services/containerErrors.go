package services

type imageError struct {
	error
	Message string
}

func (e imageError) Error() string {
	return e.Message
}
