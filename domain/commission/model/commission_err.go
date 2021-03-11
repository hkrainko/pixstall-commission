package model

type CommissionError int

func (e CommissionError) Error() string {
	switch e {
	case CommissionErrorNotFound:
		return "CommissionErrorNotFound"
	case CommissionErrorUnAuth:
		return "CommissionErrorUnAuth"
	case CommissionErrorStateNotAllowed:
		return "CommissionErrorStateNotAllowed"
	default:
		return "CommissionErrorUnknown"
	}
}

const (
	CommissionErrorNotFound            CommissionError = 10
	CommissionErrorUnAuth              CommissionError = 11
	CommissionErrorStateNotAllowed     CommissionError = 12
	CommissionErrorPriceInvalid        CommissionError = 13
	CommissionErrorDayNeedInvalid      CommissionError = 14
	CommissionErrorNotAllowBePrivate   CommissionError = 15
	CommissionErrorNotAllowAnonymous   CommissionError = 16
	CommissionErrorNotAllowSendMessage CommissionError = 16
	CommissionErrorUnknown             CommissionError = 99
)
