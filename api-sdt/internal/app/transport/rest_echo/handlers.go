package rest_echo

import (
	"api-sdt/internal/app/trace"
	"api-sdt/internal/domain/dtos"
	"api-sdt/internal/domain/entities"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func (a App) HealthCheck(c echo.Context) error {
	_, span := trace.NewSpan(c.Request().Context(), "Endpoint.HealthCheck")
	defer span.End()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"health": "ok",
		"status": http.StatusOK,
	})
}

func (a App) Create(c echo.Context) error {
	_, span := trace.NewSpan(c.Request().Context(), "Endpoint.Create")
	defer span.End()

	var objRequest entities.Projeto

	if err := c.Bind(&objRequest); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, &dtos.Error{
			Message: "Requisição inválida",
			Code:    400,
			Name:    "REQ_ERR",
		})
	}

	if err := c.Validate(&objRequest); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, &dtos.Error{
			Message: err.Error(),
			Code:    400,
			Name:    "REQ_INVALID",
		})
	}

	data, err := a.projetoSvc.Create(c.Request().Context(), &objRequest)

	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusOK, data)

}

func (a App) Update(c echo.Context) error {
	_, span := trace.NewSpan(c.Request().Context(), "Endpoint.Update")
	defer span.End()

	nome := c.Param("nome")

	if nome == "" {
		return c.JSON(http.StatusBadRequest, &dtos.Error{
			Message: "Nome inválido",
			Code:    400,
			Name:    "REQ_INVALID",
		})
	}

	var prjRequest entities.Projeto
	if err := c.Bind(&prjRequest); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, &dtos.Error{
			Message: "Requisição inválida!",
			Code:    400,
			Name:    "REQ_ERR",
		})
	}
	if err := c.Validate(&prjRequest); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, &dtos.Error{
			Message: err.Error(),
			Code:    400,
			Name:    "REQ_INVALID",
		})
	}
	data, err := a.projetoSvc.Update(c.Request().Context(), nome, &prjRequest)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, &dtos.Error{
			Message: err.Error.Error(),
			Code:    406,
			Name:    "NOT_ACCEPTED",
		})
	}
	return c.JSON(http.StatusOK, data)
}

func (a App) FindByName(c echo.Context) error {
	_, span := trace.NewSpan(c.Request().Context(), "Endpoint.FindByName")
	defer span.End()

	nome := c.Param("nome")

	if nome == "" {
		return c.JSON(http.StatusBadRequest, &dtos.Error{
			Message: "Nome inválido",
			Code:    400,
			Name:    "REQ_INVALID",
		})
	}

	data, err := a.projetoSvc.FindByName(c.Request().Context(), nome)

	if err != nil {
		return c.JSON(http.StatusNotFound, &dtos.Error{
			Message: err.Error.Error(),
			Code:    404,
			Name:    "NOT_FOUND",
		})
	}

	return c.JSON(http.StatusOK, data)
}

func (a App) FindAll(c echo.Context) error {
	_, span := trace.NewSpan(c.Request().Context(), "Endpoint.FindAll")
	defer span.End()

	data, err := a.projetoSvc.FindAll(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusNotFound, &dtos.Error{
			Message: err.Message,
			Code:    404,
			Name:    "NOT_FOUND",
		})
	}

	return c.JSON(http.StatusOK, data)

}

func (a App) Delete(c echo.Context) error {
	_, span := trace.NewSpan(c.Request().Context(), "Endpoint.Delete")
	defer span.End()

	nome := c.Param("nome")

	if nome == "" {
		return c.JSON(http.StatusBadRequest, &dtos.Error{
			Message: "Nome inválido",
			Code:    400,
			Name:    "REQ_INVALID",
		})
	}

	err := a.projetoSvc.Delete(c.Request().Context(), nome)

	if err != nil {
		return c.JSON(err.Code, &dtos.Error{
			Message: err.Error.Error(),
			Code:    err.Code,
			Name:    err.Name,
		})
	}
	return c.JSON(http.StatusOK, "Deleted")
}
