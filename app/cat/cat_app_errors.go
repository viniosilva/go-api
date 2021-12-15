package cat

type error interface {
	Error() string
}

type InvalidCmdError struct{}

func (e *InvalidCmdError) Error() string {
	return "InvalidCmdError"
}
