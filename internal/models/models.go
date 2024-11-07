package models

import (
	"regexp"
)

type ProductCategory struct {
	Name      string
	MinMonths int
	MaxMonths int
	Rate      float64
}

var ProductCategories = map[string]ProductCategory{
	"Smartphone": {Name: "Smartphone", MinMonths: 3, MaxMonths: 9},
	"Computer":   {Name: "Computer", MinMonths: 3, MaxMonths: 12},
	"TV":         {Name: "TV", MinMonths: 3, MaxMonths: 18},
}

type PaymentParams struct {
	Product     string  `json:"product"`
	Price       float64 `json:"price"`
	PhoneNumber string  `json:"phone_number"`
	Months      int     `json:"months"`
}

func (p PaymentParams) Validate() error {
	if p.Product == "" {
		return ErrProductNameIsRequired
	}

	if p.Price <= 0 {
		return ErrPriceIsRequired
	}

	phoneRegex := `^\+992\d{9}$`
	matched, _ := regexp.MatchString(phoneRegex, p.PhoneNumber)
	if !matched {
		return ErrInvalidPhoneNumberFormat
	}

	_, exists := ProductCategories[p.Product]
	if !exists {
		return ErrUnknownProductCategory
	}

	validMonths := []int{3, 6, 9, 12, 18, 24}
	isValid := false
	for _, validMonth := range validMonths {
		if validMonth == p.Months {
			isValid = true
			break
		}
	}

	if !isValid {
		return ErrInvalidMonths
	}

	return nil
}

type PaymentPlan struct {
	Months     int
	Percentage float64
}

var productPlans = map[string][]PaymentPlan{
	"Smartphone": {
		{Months: 9, Percentage: 0.00}, 
		{Months: 12, Percentage: 0.03}, 
		{Months: 18, Percentage: 0.06}, 
		{Months: 24, Percentage: 0.09},
	},
	"Computer": {
		{Months: 12, Percentage: 0.00}, 
		{Months: 18, Percentage: 0.04}, 
		{Months: 24, Percentage: 0.08}, 
	},
	"TV": {
		{Months: 18, Percentage: 0.00}, 
		{Months: 24, Percentage: 0.05}, 
	},
}

func (p PaymentParams) TotalPrice() (float64, error) {
	plans, exists := productPlans[p.Product]
	if !exists {
		return 0, ErrUnknownProductPlan
	}

	var percentage float64
	for _, plan := range plans {
		if p.Months <= plan.Months {
			percentage = plan.Percentage
			break
		}
	}

	additionalAmount := p.Price * percentage
	totalPrice := p.Price + additionalAmount

	return totalPrice, nil
}