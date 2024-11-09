package main

import (
	"company-management-api/config"
	"company-management-api/db"
	"company-management-api/handlers"
	"company-management-api/middleware"
	"company-management-api/models"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SeedAdminUser() {
    var count int64
    db.DB.Model(&models.User{}).Where("role = ?", "Admin").Count(&count)
    
    if count == 0 {
        password := "adminpassword" // Replace with your desired password
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            log.Fatal("Failed to hash password:", err)
        }

        admin := models.User{
            Username: "admin",
            Password: string(hashedPassword),
            Role:     "Admin",
        }

        db.DB.Create(&admin)
        log.Println("Admin user created with username 'admin' and provided password.")
    }
}


func main() {
	config.LoadEnv()
    // Initialize database connection
    db.Init()
    SeedAdminUser()

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
