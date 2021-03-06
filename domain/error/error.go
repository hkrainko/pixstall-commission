package error

import "strconv"

type DomainError int

func (e DomainError) Error() string {
	switch e {
	case UnknownError:
		return "UnknownError"
	default:
		return strconv.Itoa(int(e))
	}
}

const (
	UnknownError DomainError = 10
)