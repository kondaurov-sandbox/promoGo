package lib

import "strings"

type Promotion struct {
	ID             string `json:"id"`
	Price          string `json:"price"`
	ExpirationDate string `json:"expiration_date"`
}

func GetPromotionFromString(source string) Promotion {
	parts := strings.Split(source, ",")
	return Promotion{
		ID:             parts[0],
		Price:          parts[1],
		ExpirationDate: parts[2],
	}
}
