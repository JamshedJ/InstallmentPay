package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initHandler() *gin.Engine {
	router := gin.Default()
	return router
}

func Run(addr string) error {
	httpServer := &http.Server{
		Addr:         addr,
		Handler:      initHandler(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := httpServer.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
