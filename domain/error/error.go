package error

import "strconv"

type DomainError int

func (e DomainError) Error() string {
	switch e {
	case NotFoundError:
		return "NotFoundError"
	case UnAuthError:
		return "UnAuthError"
	case UnknownError:
		return "UnknownError"
	default:
		return strconv.Itoa(int(e))
	}
}

const (
	NotFoundError DomainError = 10
	UnAuthError DomainError = 11
	UnknownError DomainError = 99
)