package msg

import "pixstall-commission/domain/message/model"

type CommissionMessage struct {
	Message model.Message `json:",inline"`
}
