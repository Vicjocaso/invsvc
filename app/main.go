package main

import (
	"fmt"
	"invsvc/handlers"
	database "invsvc/lib"
	"invsvc/repos"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	db_url := os.Getenv("DSN")
	if db_url == "" {
		panic("DB_URL is required")
	}

	db, err := database.InitDB(db_url)
	if err != nil {
		log.Fatalf("couild not initialize db: %+v", err)
	}

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogLatency: true,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			value, _ := c.Get("customValueFromContext").(int)
			fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v, latency: %v\n", v.URI, v.Status, value, v.Latency.Milliseconds())
			return nil
		},
	}))

	product := e.Group("/product")

	prodcutsRepo := repos.NewProducRepo(db)
	productHandler := handlers.NewProductHandler(prodcutsRepo)

	product.GET("/", productHandler.GetAll)
	product.GET("/:id", productHandler.GetBy)
	product.POST("/", productHandler.Create)
	e.Logger.Fatal(e.Start(":1323"))

}
