package models

type ProductCategory struct {
	Name   string
	Months []int
	Rate   float64
}

var ProductCategories = map[string]ProductCategory{
	"Smartphone": {Name: "Smartphone", Months: []int{3, 6, 9}, Rate: 0.03},
	"Computer":   {Name: "Computer", Months: []int{3, 6, 9, 12}, Rate: 0.04},
	"TV":         {Name: "TV", Months: []int{3, 6, 9, 12, 18}, Rate: 0.05},
}

type PaymentParams struct {
	Product     string  `json:"product"`
	Price       float64 `json:"price"`
	PhoneNumber string  `json:"phone_number"`
	Months      int     `json:"months"`
}
