package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	db "github.com/MochammadQemalFirza/OnlineShop/config"
	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
) 

func main() {

	envPath := ".env"

	err := godotenv.Load(envPath)
	if err != nil {
		log.Println(err.Error())
		return
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	db.CreateCon()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port tidak ditemukan")
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}