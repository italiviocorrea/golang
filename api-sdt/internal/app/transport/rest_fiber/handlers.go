package rest_fiber

import (
	"api-sdt/internal/domain/dtos"
	"api-sdt/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
	//"go.opentelemetry.io/otel"

	//"go.opentelemetry.io/otel/attribute"
	"net/http"
)

func (a App) HealthCheck(c *fiber.Ctx) error {
	_, span := a.tracer.Start(c.UserContext(), "healthHandler")
	defer span.End()

	return c.Status(http.StatusOK).
		JSON(map[string]interface{}{
			"health": "ok",
			"status": http.StatusOK,
		})
}

func (a *App) FindAll(c *fiber.Ctx) error {
	_, span := a.tracer.Start(c.UserContext(), "FindAllHandler")
	defer span.End()

	data, err := a.projetoSvc.FindAll()

	return response(data, http.StatusOK, err, c)
}

func (a *App) FindByName(c *fiber.Ctx) error {
	_, span := a.tracer.Start(c.UserContext(), "FindByNameHandler")
	defer span.End()

	nome := c.Params("nome")
	if nome == "" {
		return response(nil, http.StatusBadRequest,
			&dtos.Error{
				Message: "O nome do projeto a ser pesquisado, não foi informado!",
				Code:    http.StatusBadRequest},
			c)
	}

	data, err := a.projetoSvc.FindByName(nome)
	return response(data, http.StatusOK, err, c)
}

func (a *App) Create(c *fiber.Ctx) error {
	_, span := a.tracer.Start(c.UserContext(), "CreateHandler")
	defer span.End()

	var objRequest entities.Projeto

	if err := c.BodyParser(&objRequest); err != nil {
		return response(objRequest, http.StatusBadRequest,
			&dtos.Error{
				Message: err.Error(),
				Code:    http.StatusBadRequest},
			c)
	}

	data, err := a.projetoSvc.Create(&objRequest)

	return response(data, http.StatusOK, err, c)

}

func (a *App) Update(c *fiber.Ctx) error {
	_, span := a.tracer.Start(c.UserContext(), "UpdateHandler")
	defer span.End()

	nome := c.Params("nome")

	if nome == "" {
		return response(nil, http.StatusBadRequest,
			&dtos.Error{
				Message: "O nome do projeto a ser atualizado, não foi informado!",
				Code:    http.StatusBadRequest},
			c)
	}

	var prjRequest entities.Projeto
	if err := c.BodyParser(&prjRequest); err != nil {
		return response(prjRequest, http.StatusBadRequest,
			&dtos.Error{
				Message: err.Error(),
				Code:    http.StatusBadRequest},
			c)
	}

	data, err := a.projetoSvc.Update(nome, &prjRequest)

	return response(data, http.StatusOK, err, c)
}

func (a *App) Delete(c *fiber.Ctx) error {
	_, span := a.tracer.Start(c.UserContext(), "DeleteHandler")
	defer span.End()

	nome := c.Params("nome")

	if nome == "" {
		return response(nil, http.StatusBadRequest,
			&dtos.Error{
				Message: "O nome do projeto a ser excluído, não foi informado!",
				Code:    http.StatusBadRequest},
			c)
	}

	err := a.projetoSvc.Delete(nome)

	return response(nil, http.StatusNoContent, err, c)
}

func response(data interface{}, httpStatus int, err *dtos.Error, c *fiber.Ctx) error {
	if err != nil {
		return c.Status(err.Code).JSON(map[string]string{
			"error": err.Message,
		})
	} else {
		if data != nil {
			return c.Status(httpStatus).JSON(data)
		} else {
			c.Status(httpStatus)
			return nil
		}
	}
}
