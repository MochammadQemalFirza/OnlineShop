package main

import (
	"fmt"
	"log"
	"os"

	db "github.com/MochammadQemalFirza/OnlineShop/config"
	"github.com/MochammadQemalFirza/OnlineShop/routes"
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
	db := db.CreateCon()
	e := echo.New()
	
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port tidak ditemukan")
	}

    routes.InitRouter(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}