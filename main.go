package main

import (
	"payment_service/repositories"
	"payment_service/resolvers"
	"payment_service/services"
	"payment_service/graph"
	"payment_service/helpers"
	"payment_service/handlers"

	"log"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
        log.Fatal("Error loading .env file")
    }
	
	db, err := helpers.GetGormDB()
    if err != nil {
        log.Fatal("Failed to connect to database: " + err.Error())
    }
    paymentRepository := repositories.NewPaymentRepository(db)
    paymentService := services.NewPaymentService(paymentRepository)
    resolver := resolvers.NewPaymentResolver(paymentService)

    mutationType := schema.NewMutationType(resolver)
	queryType := schema.NewQueryType(resolver)

	schema.InitSchema(queryType, mutationType)
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Add any origins you need
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))
	e.POST("/graphql", handlers.Handler)
	e.POST("/mkReq", func(c echo.Context) error {
        // We call the function from services
        // ctx is needed for the HTTP request inside SendMKRequest
        result, err := services.SendMKRequest(c.Request().Context())
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]interface{}{
                "status": "error",
                "message": err.Error(),
            })
        }
        return c.JSON(http.StatusOK, result)
    })
	e.Logger.Fatal(e.Start(":8096"))
}



