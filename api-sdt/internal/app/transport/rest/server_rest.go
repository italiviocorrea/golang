package rest

import (
	"api-sdt/internal/app/config"
	"api-sdt/internal/app/database/mongodb"
	"api-sdt/internal/domain/ports"
	"api-sdt/internal/domain/usecases"
	"api-sdt/internal/domain/usecases/validators"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
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
	log.Info("Configurando Rotas")

	apiV1 := a.server.Group("/api/v1")
	apiV1.Static("/swagger", "api/swaggerui")
	apiV1.GET("/healthy", a.HealthCheck)
	apiV1.GET("/projeto", a.FindAll)
	apiV1.POST("/projeto", a.Create)
	apiV1.GET("/projeto/:nome", a.FindByName)
	apiV1.PUT("/projeto/:nome", a.Update)
	apiV1.DELETE("/projeto/:nome", a.Delete)

}

func (a App) Start() {
	a.ConfigureRoutes()
	a.server.Start(a.cfg.SrvHost + ":" + a.cfg.SrvPort)
}
