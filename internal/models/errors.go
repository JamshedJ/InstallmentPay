package models

import "errors"

var (
	ErrValidationFailed         = errors.New("ErrValidationFailed")
	ErrProductNameIsRequired    = errors.New("ErrProductNameIsRequired")
	ErrPriceIsRequired          = errors.New("ErrPriceIsRequired")
	ErrInvalidPhoneNumberFormat = errors.New("ErrInvalidPhoneNumberFormat")
	ErrUnknownProductCategory   = errors.New("ErrUnknownProductCategory")
	ErrUnknownProductPlan       = errors.New("ErrUnknownProductPlan")
	ErrInvalidMonths            = errors.New("ErrInvalidMonths")
)
