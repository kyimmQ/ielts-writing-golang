package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/kyimmQ/ielts-writing-golang/global"
	"github.com/kyimmQ/ielts-writing-golang/internal/innitialize"
	"github.com/kyimmQ/ielts-writing-golang/internal/routes"
	"github.com/kyimmQ/ielts-writing-golang/pkg/logger"
)

type Server struct {
	port int
}

func InitServer() *http.Server {
	innitialize.LoadConfig()
	innitialize.InitMongoDB()

	port, _ := strconv.Atoi(global.Config.Server.Port)

	loggerOpts := &slog.HandlerOptions{
		AddSource: true,
	}

	newLogger := slog.New(logger.NewHandler(loggerOpts))
	slog.SetDefault(newLogger)

	newServer := &Server{
		port: port,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", newServer.port),
		Handler: routes.InitRoute(),
	}

	return server

}
