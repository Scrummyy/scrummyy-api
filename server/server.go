package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Scrummyy/scrummyy-api/internal/constants"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Serve(r *gin.Engine, conf *viper.Viper) {
	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", conf.GetString(constants.AppPort)),
		Handler:      r,
		ReadTimeout:  time.Duration(conf.GetInt(constants.ApiReadTimeout)) * time.Second,
		WriteTimeout: time.Duration(conf.GetInt(constants.ApiWriteTimeout)) * time.Second,
	}

	logrus.Infof("Ready to serve request at port %s", conf.GetString(constants.AppPort))
	go func() {
		err := s.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("HTTP Server failure: %s", err)
		}
	}()

	// r.Use(middleware.CORSMiddleware())
	// r.Use(middleware.RequestIDMiddleware())
	// r.Use(gzip.Gzip(gzip.DefaultCompression))

	// graceful shutdown code
	quit := make(chan os.Signal)
	//nolint //because of false positive
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("Shutdown initiated...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Forceful shutdown:", err)
	}

	logrus.Info("Exiting...")
}
