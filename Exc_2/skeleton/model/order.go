package model

import "time"

type Order struct {
	DrinkID uint64 `json:"drink_id"` // foreign key
	//  Add fields: CreatedAt (time.Time), Amount with suitable types
	//  json attributes need to be snakecase, i.e. name, created_at, my_variable, ..
	CreatedAt time.Time `json:"created_at"`
	Amount uint32`json:"amount"`
}
