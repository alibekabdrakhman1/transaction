package models

import "time"

type Transaction struct {
	ID                string    `json:"id"`
	Username          string    `json:"username"`
	TypeOfTransaction string    `json:"type_of_transaction"`
	Amount            float32   `json:"amount"`
	PaymentTime       time.Time `json:"payment_time"`
	Description       string    `json:"description"`
}

//var Types []string = {
//	"expenditure",
//	"replenishment"
//}
