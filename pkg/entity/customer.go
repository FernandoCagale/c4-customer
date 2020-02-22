package entity

import "github.com/lib/pq"

type Customer struct {
	Code   string         `json:"code" gorm:"primary_key"`
	Name   string         `json:"name"`
	Email  string         `json:"email"`
	Phone  string         `json:"phone"`
	Notify pq.StringArray `json:"notify" gorm:"type:varchar(64)[]"`
}
