package model

type Status int

const (
	Dead Status = iota
	Alive
)

type Gender int

const (
	Male Gender = iota
	Female
	Other
)

type Player struct {
	Number int64  `json:"number"`
	Status Status `json:"status"`
	Name   string `json:"name"`
	Gender Gender `json:"gender"`
}
