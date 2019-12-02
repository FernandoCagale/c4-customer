package entity

import "github.com/lib/pq"

type Customer struct {
	Code   string         `json:"code"`
	Name   string         `json:"name"`
	Email  string         `json:"orders"`
	Phone  string         `json:"phone"`
	Notify pq.StringArray `json:"notify",gorm:"type:varchar(64)[]"`
}
