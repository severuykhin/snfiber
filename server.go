package snfiber

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/severuykhin/logfmt/v4"
)

type config struct {
	withMetricsRoute bool
	logger           logger
}

type serverOptFunc func(cfg *config) *config

func WithMetricsRoute() serverOptFunc {
	return func(cfg *config) *config {
		cfg.withMetricsRoute = true
		return cfg
	}
}

func WithLogger(l logger) serverOptFunc {
	return func(cfg *config) *config {
		cfg.logger = l
		return cfg
	}
}

func NewServer(
	router *Router,
	opts ...serverOptFunc,
) *fiber.App {

	cfg := config{}
	for _, optFunc := range opts {
		optFunc(&cfg)
	}

	app := fiber.New()
	handler := handler{}

	if cfg.withMetricsRoute {
		app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	}

	var logger logger
	if cfg.logger != nil {
		logger = cfg.logger
	} else {
		logger = logfmt.New(os.Stdout, logfmt.L_INFO)
	}

	handler.logger = logger

	for _, route := range router.routes {
		app.Add(route.Method, string(route.Path), handler.handle(route.Handler))
	}

	return app
}
