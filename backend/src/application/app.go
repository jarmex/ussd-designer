package application

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/redis/go-redis/v9"
)

type App struct {
	app    *fiber.App
	config Config
	rdb    *redis.Client
}

func New(config Config) *App {
	return &App{
		app:    fiber.New(),
		config: config,
		rdb: redis.NewClient(&redis.Options{
			Addr: config.RedisAddress,
		}),
	}
}

func (a *App) Start(ctx context.Context) error {
	if err := a.rdb.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis client:", err)
		}
	}()

	a.app.Use(logger.New())
	a.registerRoutes(a.app.Group("/api"))

	ch := make(chan error, 1)

	go func() {
		if err := a.app.Listen(fmt.Sprintf(":%d", a.config.ServerPort)); err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		return a.app.Shutdown()
	}
}
