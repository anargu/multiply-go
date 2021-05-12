package multiply

func GetErrorMessage(err error) *string {
	if err == nil {
		return nil
	}
	msg := err.Error()
	return &msg
}
