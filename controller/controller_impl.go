package controller

import (
	"net/http"

	"github.com/MochammadQemalFirza/OnlineShop/model/web"
	service "github.com/MochammadQemalFirza/OnlineShop/service"

	"github.com/labstack/echo/v4"
)

type ControllerImpl struct {
	Service service.Service
}

func (controller *ControllerImpl) GetListProduct(c echo.Context) error {
	
	category := c.QueryParam("category")

	result, err := controller.Service.GetListProduct(c.Request().Context(), category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
	}

	if len(result) == 0 {
		return c.JSON(http.StatusNotFound, web.WebResponse{
			Message: "Data tidak tersedia!",
			Status:  http.StatusNotFound,
		})
	}

	return c.JSON(http.StatusOK, web.WebResponseSuccesForGet{
		Message: "Data berhasil didapatkan!",
		Status:  http.StatusOK,
		Data:    result,
	})
}

func NewController(Service service.Service) Controller {
	return &ControllerImpl{Service: Service}
}