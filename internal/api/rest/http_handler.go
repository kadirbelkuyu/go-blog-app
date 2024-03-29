package rest

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go-blog-app/config"
	"go-blog-app/internal/helper"
)

type RestHandler struct {
	App    *fiber.App
	DB     *sql.DB
	Auth   helper.Auth
	Redis  *redis.Client
	Config config.AppConfig
}
