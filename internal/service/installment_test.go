package service

import (
	"testing"

	"github.com/JamshedJ/InstallmentPay/internal/models"
	"github.com/stretchr/testify/require"
)

func TestCalculatePayment(t *testing.T) {
	tests := []struct {
		testCase       string
		params         models.PaymentParams
		err            error
		expectedAmount float64
	}{
		{
			testCase: "Valid Smartphone Installment",
			params: models.PaymentParams{
				Product:     "Smartphone",
				Price:       1000,
				PhoneNumber: "+992123456789",
				Months:      9,
			},
			err:            nil,
			expectedAmount: 1000,
		},
		{
			testCase: "Valid Computer Installment with Percentage",
			params: models.PaymentParams{
				Product:     "Computer",
				Price:       1500,
				PhoneNumber: "+992987654321",
				Months:      18,
			},
			err:            nil,
			expectedAmount: 1560,
		},
		{
			testCase: "Empty product name",
			params: models.PaymentParams{
				Product:     "",
				Price:       1000,
				PhoneNumber: "+992123456789",
				Months:      6,
			},
			err:            models.ErrProductNameIsRequired,
			expectedAmount: 0,
		},
		{
			testCase: "Price is required",
			params: models.PaymentParams{
				Product:     "Smartphone",
				Price:       0,
				PhoneNumber: "+992123456789",
				Months:      6,
			},
			err:            models.ErrPriceIsRequired,
			expectedAmount: 0,
		},
		{
			testCase: "Invalid phone number format",
			params: models.PaymentParams{
				Product:     "Smartphone",
				Price:       1000,
				PhoneNumber: "123456789",
				Months:      6,
			},
			err:            models.ErrInvalidPhoneNumberFormat,
			expectedAmount: 0,
		},
		{
			testCase: "Invalid product category",
			params: models.PaymentParams{
				Product:     "UnknownProduct",
				Price:       1000,
				PhoneNumber: "+992123456789",
				Months:      6,
			},
			err:            models.ErrUnknownProductCategory,
			expectedAmount: 0,
		},
		{
			testCase: "Invalid months",
			params: models.PaymentParams{
				Product:     "Smartphone",
				Price:       1000,
				PhoneNumber: "+992123456789",
				Months:      30,
			},
			err:            models.ErrInvalidMonths,
			expectedAmount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			totalAmount, err := CalculatePayment(tt.params)

			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				require.Equal(t, 0.0, totalAmount)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedAmount, totalAmount)
			}
		})
	}
}
