package error

import "strconv"

type DomainError int

func (e DomainError) Error() string {
	switch e {
	case UnknownError:
		return "UnknownError"
	case UnAuthError:
		return "UnAuthError"
	default:
		return strconv.Itoa(int(e))
	}
}

const (
	NotFoundError DomainError = 10
	UnAuthError DomainError = 11
	UnknownError DomainError = 99
)