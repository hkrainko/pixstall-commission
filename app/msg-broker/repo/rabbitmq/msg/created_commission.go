package msg

import "pixstall-commission/domain/commission/model"

type CreatedCommission struct {
	Comm model.Commission `json:",inline"`
}