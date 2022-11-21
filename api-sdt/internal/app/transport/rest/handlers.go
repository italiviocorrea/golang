package rest

import (
	"api-sdt/internal/domain/dtos"
	"api-sdt/internal/domain/entities"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func (a App) HealthCheck(c echo.Context) error {
	log.Info("HealthCheck")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "O servidor está funcionando",
	})
}

func (a App) Create(c echo.Context) error {

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

	data, err := a.projetoSvc.Create(&objRequest)

	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusOK, data)

}

func (a App) Update(c echo.Context) error {

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
	data, err := a.projetoSvc.Update(nome, &prjRequest)
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

	nome := c.Param("nome")

	if nome == "" {
		return c.JSON(http.StatusBadRequest, &dtos.Error{
			Message: "Nome inválido",
			Code:    400,
			Name:    "REQ_INVALID",
		})
	}

	data, err := a.projetoSvc.FindByName(nome)

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

	data, err := a.projetoSvc.FindAll()

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

	nome := c.Param("nome")

	if nome == "" {
		return c.JSON(http.StatusBadRequest, &dtos.Error{
			Message: "Nome inválido",
			Code:    400,
			Name:    "REQ_INVALID",
		})
	}

	err := a.projetoSvc.Delete(nome)

	if err != nil {
		return c.JSON(err.Code, &dtos.Error{
			Message: err.Error.Error(),
			Code:    err.Code,
			Name:    err.Name,
		})
	}
	return c.JSON(http.StatusOK, "Deleted")
}
