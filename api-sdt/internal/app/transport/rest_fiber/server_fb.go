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
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel"

	//stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/exporters/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type App struct {
	server     *fiber.App
	projetoSvc ports.ProjetoUseCasePort
	cfg        *config.Settings
	tracer     trace.Tracer
}

var tracer = otel.Tracer("API-SDT_FIBER")

func New(cfg *config.Settings, client *mongo.Client) *App {

	tp, tpErr := JaegerTraceProvider(cfg)

	if tpErr != nil {
		log.Fatal(tpErr)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	//defer func() {
	//	if err := tp.Shutdown(context.Background()); err != nil {
	//		log.Printf("Error shutting down tracer provider: %v", err)
	//	}
	//}()

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
		tracer:     tracer,
	}

}

func (a App) ConfigureRoutes() {
	log.Info("Configurando Rotas")

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

func JaegerTraceProvider(cfg *config.Settings) (*sdktrace.TracerProvider, error) {
	//exporter, err := stdout.New(stdout.WithPrettyPrint())
	log.Info(cfg.JaegerEndpoint)
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(cfg.JaegerEndpoint)))
	if err != nil {
		log.Fatal(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("API-SDT"),
				semconv.DeploymentEnvironmentKey.String(cfg.Env),
			)),
	)
	return tp, nil
}
