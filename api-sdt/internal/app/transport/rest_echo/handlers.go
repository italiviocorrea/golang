package rest_echo

import (
	"api-sdt/internal/app/trace"
	"api-sdt/internal/domain/dtos"
	"api-sdt/internal/domain/entities"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a App) HealthCheck(c echo.Context) error {
	_, span := trace.NewSpan(c.Request().Context(), "ProjetoEndpoint.HealthCheck")
	defer span.End()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"health": "ok",
		"status": http.StatusOK,
	})
}

func (a App) Create(c echo.Context) error {
	_, span := trace.NewSpan(c.Request().Context(), "ProjetoEndpoint.Create")
	defer span.End()

	var objRequest entities.Projeto

	if err := c.Bind(&objRequest); err != nil {
		return Response(objRequest, http.StatusBadRequest,
			&dtos.Error{
				Message: err.Error(),
				Code:    http.StatusBadRequest},
			c)
	}

	if err := c.Validate(&objRequest); err != nil {
		return Response(objRequest, http.StatusBadRequest,
			&dtos.Error{
				Message: err.Error(),
				Code:    http.StatusBadRequest},
			c)
	}

	data, err := a.projetoSvc.Create(c.Request().Context(), &objRequest)

	return Response(data, http.StatusOK, err, c)

}

func (a App) Update(c echo.Context) error {
	_, span := trace.NewSpan(c.Request().Context(), "ProjetoEndpoint.Update")
	defer span.End()

	nome := c.Param("nome")

	if nome == "" {
		return Response(nil, http.StatusBadRequest,
			&dtos.Error{
				Message: "O nome do projeto a ser atualizado, não foi informado!",
				Code:    http.StatusBadRequest},
			c)
	}

	var prjRequest entities.Projeto

	if err := c.Bind(&prjRequest); err != nil {
		return Response(prjRequest, http.StatusBadRequest,
			&dtos.Error{
				Message: err.Error(),
				Code:    http.StatusBadRequest},
			c)
	}

	if err := c.Validate(&prjRequest); err != nil {
		return Response(prjRequest, http.StatusBadRequest,
			&dtos.Error{
				Message: err.Error(),
				Code:    http.StatusBadRequest},
			c)
	}

	data, err := a.projetoSvc.Update(c.Request().Context(), nome, &prjRequest)

	return Response(data, http.StatusOK, err, c)
}

func (a App) FindByName(c echo.Context) error {
	_, span := trace.NewSpan(c.Request().Context(), "ProjetoEndpoint.FindByName")
	defer span.End()

	nome := c.Param("nome")

	if nome == "" {
		return Response(nil, http.StatusBadRequest,
			&dtos.Error{
				Message: "O nome do projeto a ser pesquisado, não foi informado!",
				Code:    http.StatusBadRequest},
			c)
	}

	data, err := a.projetoSvc.FindByName(c.Request().Context(), nome)

	return Response(data, http.StatusOK, err, c)
}

func (a App) FindAll(c echo.Context) error {
	_, span := trace.NewSpan(c.Request().Context(), "ProjetoEndpoint.FindAll")
	defer span.End()

	data, err := a.projetoSvc.FindAll(c.Request().Context())

	return Response(data, http.StatusOK, err, c)

}

func (a App) Delete(c echo.Context) error {
	_, span := trace.NewSpan(c.Request().Context(), "ProjetoEndpoint.Delete")
	defer span.End()

	nome := c.Param("nome")

	if nome == "" {
		return Response(nil, http.StatusBadRequest,
			&dtos.Error{
				Message: "O nome do projeto a ser excluído, não foi informado!",
				Code:    http.StatusBadRequest},
			c)
	}

	err := a.projetoSvc.Delete(c.Request().Context(), nome)

	return Response(nil, http.StatusNoContent, err, c)
}

func Response(data interface{}, httpStatus int, err *dtos.Error, c echo.Context) error {
	if err != nil {
		return c.JSON(httpStatus, map[string]string{
			"error": err.Message,
		})
	} else {
		if data != nil {
			return c.JSON(httpStatus, data)
		} else {
			c.JSON(httpStatus, "")
			return nil
		}
	}
}
