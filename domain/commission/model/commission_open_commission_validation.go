package model

type CommissionOpenCommissionValidation struct {
	ID                             string  `json:"id"`
	IsValid                        bool    `json:"isValid"`
	InvalidReason                  *string `json:"invalidReason,omitempty"`
	TimesAllowedDraftToChange      *int    `json:"timesAllowedDraftToChange,omitempty"`
	TimesAllowedCompletionToChange *int    `json:"timesAllowedCompletionToChange,omitempty"`
}
