package main

import (
	"github.com/JRKs1532/sa-65-example/controller"
	"github.com/JRKs1532/sa-65-example/entity"
	"github.com/JRKs1532/sa-65-example/middlewares"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{

		router.Use(middlewares.Authorizes())
		{
			// Employee Routes
			router.GET("/employees", controller.ListEmployees)
			//router.GET("/employee/:id", controller.ListEmployees)
			router.GET("/employee/:id", controller.GetEmployee)
			//router.POST("/employees", controller.CreateEmployee)
			router.PATCH("/employees", controller.UpdateEmployee)
			router.DELETE("/employees/:id", controller.DeleteEmployee)

			// gander Routes
			router.GET("/genders", controller.ListGenders)
			router.GET("/gender/:id", controller.GetGender)
			router.POST("/genders", controller.CreateGender)
			router.PATCH("/genders", controller.UpdateGender)
			router.DELETE("/genders/:id", controller.DeleteGender)

			// provinces Routes
			router.GET("/provinces", controller.ListProvinces)
			router.GET("/province/:id", controller.GetProvince)
			router.POST("/provinces", controller.CreateProvince)
			router.PATCH("/provinces", controller.UpdateProvince)
			router.DELETE("/provinces/:id", controller.DeleteProvince)

			// Member Routes
			router.GET("/members", controller.ListMembers)
			router.GET("/member/:id", controller.GetMember)
			router.POST("/members", controller.CreateMember)
			router.PATCH("/members", controller.UpdateMember)
			router.DELETE("/members/:id", controller.DeleteMember)

			// Typeproduct Routes
			router.GET("/typeproducts", controller.ListTypeproducts)
			router.GET("/typeproduct/:id", controller.GetTypeproduct)
			router.POST("/typeproducts", controller.CreateTypeproduct)
			router.PATCH("/typeproducts", controller.UpdateTypeproduct)
			router.DELETE("/typeproducts/:id", controller.DeleteTypeproduct)

			// Manufacturer Routes
			router.GET("/manufacturers", controller.ListManufacturers)
			router.GET("/manufacturer/:id", controller.GetManufacturer)
			router.POST("/manufacturers", controller.CreateManufacturer)
			router.PATCH("/manufacturers", controller.UpdateManufacturer)
			router.DELETE("/manufacturers/:id", controller.DeleteManufacturer)

			// Product Routes
			router.GET("/products", controller.ListProducts)
			router.GET("/product/:id", controller.GetProduct)
			router.POST("products", controller.CreateProduct)
			router.PATCH("/products", controller.UpdateProduct)
			router.DELETE("/products/:id", controller.DeleteProduct)
			// Run the server

		}
	}
	// Signup User Route
	r.POST("/signup", controller.CreateEmployee)
	// login User Route
	r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run("localhost: " + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
