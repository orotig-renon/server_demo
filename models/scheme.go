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

type MarkerTestResponse struct {
	OroTigTask []MarkerTest `json:"OroTigTask"`
}

type MarkerTest struct {
	OrderId     string    `json:"OrderId" gorm:"primaryKey"`
	RowNo       int32     `json:"RowNo"`
	ReciptDate  time.Time `json:"ReciptDate" gorm:"autoCreateTime"`
	FileId      string    `json:"FileId" `
	Description string    `json:"Description"`
	UserId      string    `json:"UserId"`
	Dongle      string    `json:"Dongle"`
	StartDate   time.Time `json:"StartDate"`
	PiecesToGo  int32     `json:"PiecesToGo" gorm:"default:0"`
	PiecesMade  int32     `json:"PiecesMade" gorm:"default:0"`
	Note        string    `json:"Note"`
	IsComplete  bool      `json:"IsComplete"`
	EndDate     time.Time `json:"EndDate"`
}
