package model

type Order struct {
	ID     string
	Items  []CartItem
	Total  float64
	Status string // paid, failed
}
