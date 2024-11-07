package service

import (
	"errors"
	"fmt"

	"github.com/JamshedJ/InstallmentPay/internal/models"
	"github.com/JamshedJ/InstallmentPay/pkg/glog"
)

func CalculatePayment(p models.PaymentParams) (float64, error) {
	logger := glog.NewLogger()

	if err := p.Validate(); err != nil {
		logger.Error().Err(err).Msg("Validation failed")
		return 0, errors.Join(models.ErrValidationFailed, err)
	}

	totalPrice, err := p.TotalPrice()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to calculate total price")
		return 0, errors.Join(models.ErrValidationFailed, err)
	}

	message := fmt.Sprintf(
		`Your installment payment for a %s is confirmed. Total price: %.2f. Installment plan: %d months.`, 
		p.Product, totalPrice, p.Months,
	)

	err = SendSMS(p.PhoneNumber, message)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to send SMS")
		return 0, err
	}

	return totalPrice, nil
}

func SendSMS(phoneNumber string, message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", phoneNumber, message)
	return nil
}
