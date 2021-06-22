package model

type UserState string

const (
	UserStatePending    UserState = "P" //Allow to change userID
	UserStateActive     UserState = "A"
	UserStateTerminated UserState = "T"
)