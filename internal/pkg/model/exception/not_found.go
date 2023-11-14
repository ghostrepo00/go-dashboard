package exception

type NotFound struct{}

func (err *NotFound) Error() string {
	return "Record not found"
}
