package models

import (
	"gorm.io/gorm"
)

type SalaryHistory struct {
	gorm.Model
	FromUserId string `json:"from_user_id"`
	ToUserId   string `json:"to_user_id"`
	Amount     uint   `json:"amount"`
}
