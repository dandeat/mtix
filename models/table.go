package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Tiket struct {
	Ticket_ID   uint `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Cust_ID     uint `json:"cust_id"`
	Payment_ID  uint `json:"payment_id"`
	Schedule_ID uint `json:"schedule_id"`
}

type Cust struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name string `json:"name"`
}

type Movie struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name        string `json:"name"`
	Description string
}

type Payment struct {
	ID           uint            `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Cust_ID      uint            `json:"cust_id"`
	Payment_Type string          `json:"payment_type"` //cash,emoney,debit,etc
	Total        decimal.Decimal `json:"total" sql:"type:decimal(16,2)"`
	CreatedDate  time.Time       `json:"created_date"`
}

type Schedule struct {
	ID         uint            `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Movie_ID   uint            `json:"movie_id"`
	Start_Date time.Time       `json:"start_date"`
	End_Date   time.Time       `json:"end_date"`
	Room       int             `json:"room"`
	Price      decimal.Decimal `json:"price" sql:"type:decimal(16,2)"`
}
