package main

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mr-amirfazel/chisai/db"
	"github.com/mr-amirfazel/chisai/handlers"
	"github.com/mr-amirfazel/chisai/config"
)

func main() {

	config, er := config.LoadConfig()
	if er != nil {
        fmt.Println("Error loading config: %s\n", er)
        return
    }

	e := echo.New()

    // Connect to MongoDB
    err := db.ConnectDB()
    if err != nil {
        panic(err)
    }

    // Routes
    e.POST("/shorten", handlers.Shorten_url)
    e.GET("/:shortURL", handlers.RedirectToLongURL)

    // Start server
    e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.Server.Port)))
}