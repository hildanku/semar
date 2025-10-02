package server

import (
	"github.com/gofiber/fiber/v2"

	"semar-siem/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "semar-siem",
			AppName:      "semar-siem",
		}),

		db: database.New(),
	}

	return server
}
