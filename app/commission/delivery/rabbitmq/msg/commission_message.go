package msg

import "pixstall-commission/domain/message/model"

type CommissionMessage struct {
	model.Message `json:",inline"`
}
