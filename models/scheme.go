package models

import (
	"time"
)

type WelderTest struct {
	OrderId     string    `json:"order_id" gorm:"primaryKey"`
	ReciptDate  time.Time `json:"recipt_date" gorm:"autoCreateTime"`
	Program     string    `json:"program"`
	Description string    `json:"description"`
	SN          string    `json:"sn"`
	StartDate   time.Time `json:"start_date"`
	PiecesToGo  int32     `json:"pieces_to_go" `
	PiecesMade  int32     `json:"pieces_made" gorm:"default:0"`
	Note        string    `json:"note"`
	IsComplete  bool      `json:"complete"`
	EndDate     time.Time `json:"end_date"`
}
