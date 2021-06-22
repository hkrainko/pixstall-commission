package model

import "time"

type UserUpdater struct {
	UserID      string     `json:"userId"`
	UserName    *string    `json:"userName,omitempty"`
	Email       *string    `json:"email,omitempty"`
	Birthday    *string    `json:"birthday,omitempty"`
	Gender      *string    `json:"gender,omitempty"`
	ProfilePath *string    `json:"profilePath,omitempty"`
	State       *UserState `json:"state,omitempty"`
	IsArtist    *bool      `json:"isArtist,omitempty"`
	RegTime     *time.Time `json:"regTime,omitempty"`
}