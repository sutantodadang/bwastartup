package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main()  {
	dsn := "host=localhost user=postgres password=admin dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Database Terkoneksi")

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	// userInput := user.RegisterUserInput{
	// 	Name: "sutanto",
	// 	Email: "sutanto@mail.com",
	// 	Occupation: "anak gaul",
	// 	Password: "password",
	// }

	// userService.RegisterUser(userInput)

	// user := user.User{
	// 	Id: 4,
	// 	Name: "dadang",
	// }

	// userRepo.Save(user)

	// var users []user.User
	// db.Find(&users)

	// fmt.Println("Isi slicenya : ", len(users))

	router := gin.Default()

	// router.GET("/handler",handler)
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailibity)

	router.Run()
}

// func handler(c *gin.Context)  {
// 	dsn := "host=localhost user=postgres password=admin dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	var users []user.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)

// }