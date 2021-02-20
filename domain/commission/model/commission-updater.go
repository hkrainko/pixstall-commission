package model

import "time"

type CommissionUpdater struct {
	ID                             string
	CompleteTime                   *time.Time
	State                          *CommissionState
}