package handler

import (
	"net/http"

	"github.com/JamshedJ/InstallmentPay/internal/models"
	"github.com/JamshedJ/InstallmentPay/internal/service"
	"github.com/gin-gonic/gin"
)

func CalculateInstallment(c *gin.Context) {
	var params models.PaymentParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	totalAmount, err := service.CalculatePayment(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_amount": totalAmount})
}
