package commons

type GeneralError struct {
	Message string
}

func (e GeneralError) Error() string { return e.Message }
