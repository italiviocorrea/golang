package rest_fiber

import (
	"api-sdt/internal/app/config"
	"api-sdt/internal/app/database/mongodb"
	"api-sdt/internal/domain/ports"
	"api-sdt/internal/domain/usecases"
	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.opentelemetry.io/otel/trace"
	//"go.opentelemetry.io/otel/trace"
	//"go.opentelemetry.io/otel/trace"
)

type App struct {
	server     *fiber.App
	projetoSvc ports.ProjetoUseCasePort
	cfg        *config.Settings
}

func New(cfg *config.Settings, client *mongo.Client) *App {

	server := fiber.New()
	server.Use(otelfiber.Middleware("API-SDT"))
	server.Use(logger.New())
	server.Use(cors.New())

	// injeta o repositorio
	projetoRepository := mongodb.NewRepository(cfg, client)

	// injeta os casos de usos
	projetoSvc := usecases.New(cfg, projetoRepository)

	return &App{
		server:     server,
		projetoSvc: projetoSvc,
		cfg:        cfg,
	}

}

func (a App) ConfigureRoutes() {

	apiV1 := a.server.Group("/api/v1")
	apiV1.Static("/swagger", "api/swaggerui")
	apiV1.Get("/healthy", a.HealthCheck)
	apiV1.Get("/projetos", a.FindAll)
	apiV1.Post("/projetos", a.Create)
	apiV1.Get("/projetos/:nome", a.FindByName)
	apiV1.Put("/projetos/:nome", a.Update)
	apiV1.Delete("/projetos/:nome", a.Delete)

}

func (a App) Start() {
	a.ConfigureRoutes()
	a.server.Listen(a.cfg.SrvHost + ":" + a.cfg.SrvPort)
}
