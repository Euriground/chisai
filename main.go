package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mr-amirfazel/chisai/handlers"
	"github.com/mr-amirfazel/chisai/db"

)

func main() {

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
    e.Logger.Fatal(e.Start(":8080"))
}