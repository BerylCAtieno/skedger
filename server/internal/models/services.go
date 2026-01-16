package models

import "time"

type Service struct {
	Id         string    `json:"id"`
	BusinessId string    `json:"business_id"`
	Name       string    `json:"name"`
	Duration   string    `json:"duration"`
	Price      int64     `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
}
