package models

import "time"

type StatusType string

const (
	StatusRepaired    StatusType = "repaired"
	StatusNotRepaired StatusType = "not_repaired"
)

type Car struct {
	ID               int        `json:"id"`
	Brand            string     `json:"brand"`
	Model            string     `json:"model"`
	Owner            string     `json:"owner"`
	Status           StatusType `json:"status"`
	LastDateOfRepair time.Time  `json:"date_of_repair"`
}
