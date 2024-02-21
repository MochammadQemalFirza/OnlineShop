package controller

import "github.com/labstack/echo/v4"

type Controller interface {
	// dosen
	GetListProduct(c echo.Context) error

}