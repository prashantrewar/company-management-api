package main

import (
    "github.com/gin-gonic/gin"
    "company-management-api/handlers"
    "company-management-api/db"
    "company-management-api/middleware"
	"company-management-api/config"
)

func main() {
	config.LoadEnv()
    // Initialize database connection
    db.Init()

    router := gin.Default()

    router.Static("/static", "./frontend")

    router.GET("/", func(c *gin.Context) {
        c.File("./frontend/index.html")
    })

    // Public endpoints
    router.POST("/login", handlers.Login)

    // Authenticated endpoints
    authorized := router.Group("/")
    authorized.Use(middleware.AuthMiddleware())
    {
        authorized.POST("/users", handlers.CreateUser)
        authorized.GET("/users", handlers.GetUsers)
        authorized.POST("/payroll", handlers.CreatePayroll)
        authorized.GET("/payroll", handlers.GetPayrolls)
        authorized.POST("/customers", handlers.CreateCustomer)
        authorized.GET("/customers", handlers.GetCustomers)
        authorized.POST("/billings", handlers.CreateBilling)
        authorized.GET("/billings", handlers.GetBillings)
    }

    router.Run(":8080")
}
