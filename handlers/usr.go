package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mr-amirfazel/chisai/db"
	"github.com/mr-amirfazel/chisai/models"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"context"
)

func shorten_url(c echo.Context) error{
	db := db.GetDBInstance()
	url_collection := db.Database("chisai").Collection("URL")


	var url models.URL
	if err := c.Bind(&url); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
    }
    // Generate short URL logic here
    shortURL := generateShortURL(url.LongURL)

    // Insert URL into MongoDB
    url.ShortURL = shortURL
    _, err := url_collection.InsertOne(context.Background(), url)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert URL"})
    }

    response := map[string]string{"short_url": shortURL}
    return c.JSON(http.StatusCreated, response)
}

func RedirectToLongURL(c echo.Context) error {
    db := db.GetDBInstance()
    url_collection := db.Database("chisai").Collection("URL") // Replace with your DB and collection names

    shortURL := c.Param("shortURL")

    var url models.URL
    err := url_collection.FindOne(context.Background(), bson.M{"short_url": shortURL}).Decode(&url)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return c.JSON(http.StatusNotFound, map[string]string{"error": "URL not found"})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
    }

    return c.Redirect(http.StatusFound, url.LongURL)
}

func generateShortURL(longURL string) string {
    // shortening logic

    // This is a simplistic example using the first 5 characters of the long URL
    return "shortened/" + longURL[:5]
}