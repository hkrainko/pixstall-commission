package model

type ValidatedCommission struct {
	ID      string  `json:"id"`
	IsValid bool    `json:"isValid"`
	Reason  *string `json:"reason,omitempty"`
}
