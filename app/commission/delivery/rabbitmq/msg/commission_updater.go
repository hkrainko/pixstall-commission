package msg

import "pixstall-commission/domain/commission/model"

type CommissionUpdater struct {
	Updater model.CommissionUpdater `json:",inline"`
}
