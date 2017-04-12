package main

import "time"

type BaseModel struct {
	ID uint			`gorm: "primary_key" json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

type Thing struct {
	BaseModel
	Name string		`json:"name"`
	Description string	`json:"description"`
}

type Things []Thing
