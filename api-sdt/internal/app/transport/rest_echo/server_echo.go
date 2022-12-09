package rest_echo

import (
	"api-sdt/internal/app/config"
	"api-sdt/internal/app/database/mongodb"
	"api-sdt/internal/domain/ports"
	"api-sdt/internal/domain/usecases"
	"api-sdt/internal/domain/usecases/validators"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

type App struct {
	server     *echo.Echo
	projetoSvc ports.ProjetoUseCasePort
	cfg        *config.Settings
}

func New(cfg *config.Settings, client *mongo.Client) *App {
	server := echo.New()

	server.Validator = validators.InitCustomValidator()

	server.Use(middleware.Recover())
	server.Use(otelecho.Middleware("API-SDT"))
	server.Use(middleware.Logger())
	server.Use(middleware.RequestID())
	server.Use(middleware.CORS())

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
	apiV1.GET("/healthy", a.HealthCheck)
	apiV1.GET("/projetos", a.FindAll)
	apiV1.POST("/projetos", a.Create)
	apiV1.GET("/projetos/:nome", a.FindByName)
	apiV1.PUT("/projetos/:nome", a.Update)
	apiV1.DELETE("/projetos/:nome", a.Delete)

}

func (a App) Start() {
	a.ConfigureRoutes()
	a.server.Start(a.cfg.SrvHost + ":" + a.cfg.SrvPort)
}
