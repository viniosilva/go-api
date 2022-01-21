package exception

type error interface {
	Error() string
}

type InvalidDateException struct{}

func (e *InvalidDateException) Error() string {
	return "InvalidDateException"
}
