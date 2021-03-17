package model

type CommissionError int

func (e CommissionError) Error() string {
	switch e {
	case CommissionErrorNotFound:
		return "CommissionErrorNotFound"
	case CommissionErrorUnAuth:
		return "CommissionErrorUnAuth"
	case CommissionErrorDecisionNotAllowed:
		return "CommissionErrorDecisionNotAllowed"
	case CommissionErrorPriceInvalid:
		return "CommissionErrorPriceInvalid"
	case CommissionErrorDayNeedInvalid:
		return "CommissionErrorDayNeedInvalid"
	case CommissionErrorNotAllowBePrivate:
		return "CommissionErrorNotAllowBePrivate"
	case CommissionErrorNotAllowAnonymous:
		return "CommissionErrorNotAllowAnonymous"
	case CommissionErrorNotAllowSendMessage:
		return "CommissionErrorNotAllowSendMessage"
	default:
		return "CommissionErrorUnknown"
	}
}

const (
	CommissionErrorNotFound            CommissionError = 10
	CommissionErrorUnAuth              CommissionError = 11
	CommissionErrorDecisionNotAllowed  CommissionError = 12
	CommissionErrorPriceInvalid        CommissionError = 13
	CommissionErrorDayNeedInvalid      CommissionError = 14
	CommissionErrorNotAllowBePrivate   CommissionError = 15
	CommissionErrorNotAllowAnonymous   CommissionError = 16
	CommissionErrorNotAllowSendMessage CommissionError = 17
	CommissionErrorRevisionExceed      CommissionError = 17
	CommissionErrorUnknown             CommissionError = 99
)
