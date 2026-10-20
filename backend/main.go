package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/viniciusalexandrre/oficina/backend/config"
	"github.com/viniciusalexandrre/oficina/backend/customer"
	"github.com/viniciusalexandrre/oficina/backend/database"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.LoadEnv,
			database.DatabaseConnection,
			NewFiberApp,
			customer.NewCustomerModule,
		),
		fx.Invoke(
			RegisterRoutes,
			StartServer,
		),
	).Run()
}

func NewFiberApp() *fiber.App {
	return fiber.New()
}

func StartServer(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := app.Listen(":8080"); err != nil {
					log.Println(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.ShutdownWithContext(ctx)
		},
	})
}
