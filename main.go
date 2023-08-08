package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	docs "github.com/splax-s/ecommerce-backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/splax-s/ecommerce-backend/controllers"
	"github.com/splax-s/ecommerce-backend/database"
	"github.com/splax-s/ecommerce-backend/middleware"
	"github.com/splax-s/ecommerce-backend/routes"
)

// @title Ecommerce Backend API
// @version 1.0
// @description This is the API documentation for the Ecommerce Backend.
// @host localhost:8000
// @BasePath /
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	docs.SwaggerInfo.BasePath = "/"
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.Authentication())
	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())
	log.Fatal(router.Run(":" + port))
}
