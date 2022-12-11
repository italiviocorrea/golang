package rest_fiber

import (
	"api-sdt/internal/app/trace"
	"api-sdt/internal/domain/dtos"
	"api-sdt/internal/domain/entities"
	"api-sdt/internal/domain/usecases/validators"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var validate = validators.InitCustomValidator()

func (a App) HealthCheck(c *fiber.Ctx) error {
	_, span := trace.NewSpan(c.UserContext(), "ProjetoEndpoint.HealthCheck")
	defer span.End()

	return c.Status(http.StatusOK).
		JSON(map[string]interface{}{
			"health": "ok",
			"status": http.StatusOK,
		})
}

func (a *App) FindAll(c *fiber.Ctx) error {
	_, span := trace.NewSpan(c.UserContext(), "ProjetoEndpoint.FindAll")
	defer span.End()

	data, err := a.projetoSvc.FindAll(c.UserContext())

	return Response(data, http.StatusOK, err, c)
}

func (a *App) FindByName(c *fiber.Ctx) error {
	_, span := trace.NewSpan(c.UserContext(), "ProjetoEndpoint.FindByName")
	defer span.End()

	nome := c.Params("nome")
	if nome == "" {
		return Response(nil, http.StatusBadRequest,
			&dtos.Error{
				Message: "O nome do projeto a ser pesquisado, não foi informado!",
				Code:    http.StatusBadRequest},
			c)
	}

	data, err := a.projetoSvc.FindByName(c.UserContext(), nome)
	return Response(data, http.StatusOK, err, c)
}

func (a *App) Create(c *fiber.Ctx) error {
	_, span := trace.NewSpan(c.UserContext(), "ProjetoEndpoint.Create")
	defer span.End()

	var objRequest entities.Projeto

	if err := c.BodyParser(&objRequest); err != nil {
		return Response(objRequest, http.StatusBadRequest,
			&dtos.Error{
				Message: err.Error(),
				Code:    http.StatusBadRequest},
			c)
	}

	if err := validate.Validate(&objRequest); err != nil {
		return Response(objRequest, http.StatusBadRequest,
			&dtos.Error{
				Message: err.Error(),
				Code:    http.StatusBadRequest},
			c)
	}

	data, err := a.projetoSvc.Create(c.UserContext(), &objRequest)

	return Response(data, http.StatusOK, err, c)

}

func (a *App) Update(c *fiber.Ctx) error {
	_, span := trace.NewSpan(c.UserContext(), "ProjetoEndpoint.Update")
	defer span.End()

	nome := c.Params("nome")

	if nome == "" {
		return Response(nil, http.StatusBadRequest,
			&dtos.Error{
				Message: "O nome do projeto a ser atualizado, não foi informado!",
				Code:    http.StatusBadRequest},
			c)
	}

	var prjRequest entities.Projeto
	if err := c.BodyParser(&prjRequest); err != nil {
		return Response(prjRequest, http.StatusBadRequest,
			&dtos.Error{
				Message: err.Error(),
				Code:    http.StatusBadRequest},
			c)
	}

	if err := validate.Validate(&prjRequest); err != nil {
		return Response(prjRequest, http.StatusBadRequest,
			&dtos.Error{
				Message: err.Error(),
				Code:    http.StatusBadRequest},
			c)
	}

	data, err := a.projetoSvc.Update(c.UserContext(), nome, &prjRequest)

	return Response(data, http.StatusOK, err, c)
}

func (a *App) Delete(c *fiber.Ctx) error {
	_, span := trace.NewSpan(c.UserContext(), "ProjetoEndpoint.Delete")
	defer span.End()

	nome := c.Params("nome")

	if nome == "" {
		return Response(nil, http.StatusBadRequest,
			&dtos.Error{
				Message: "O nome do projeto a ser excluído, não foi informado!",
				Code:    http.StatusBadRequest},
			c)
	}

	err := a.projetoSvc.Delete(c.UserContext(), nome)

	return Response(nil, http.StatusNoContent, err, c)
}

func Response(data interface{}, httpStatus int, err *dtos.Error, c *fiber.Ctx) error {
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
