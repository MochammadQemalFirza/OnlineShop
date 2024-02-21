package routes

import (
	"database/sql"

	"github.com/MochammadQemalFirza/OnlineShop/controller"
	repository "github.com/MochammadQemalFirza/OnlineShop/repository"
	"github.com/MochammadQemalFirza/OnlineShop/service"

	"github.com/labstack/echo/v4"
)


func InitRouter(g *echo.Echo, db *sql.DB) {
	var (
		repository = repository.NewRepository(db)
		service    = service.NewService(repository)
		controller = controller.NewController(service)
	)

	g.GET("/product", controller.GetListProduct)
}

