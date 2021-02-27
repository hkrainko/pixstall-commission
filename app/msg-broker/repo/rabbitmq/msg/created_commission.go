package msg

import "pixstall-commission/domain/commission/model"

type CreatedCommission struct {
	model.Commission `json:",inline"`
}