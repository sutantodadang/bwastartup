package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)


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
	api.POST("/avatars", authMiddleware(authService,userService), userHandler.UploadAvatar)

	router.Run()
}

func authMiddleware (authService auth.Service, userService user.Service) gin.HandlerFunc  {
	return func (c *gin.Context)  {
		authHeader := c.GetHeader("Authorization")
	
		if !strings.Contains(authHeader,"Bearer") {
			res := helper.APIResponse("UnAuthorization", http.StatusUnauthorized,"error",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,res)
			return
		}
	
		tokenString := ""
		sliceToken := strings.Split(authHeader, " ")
		if len(sliceToken) == 2 {
			tokenString = sliceToken[1]
		}
	
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			res := helper.APIResponse("UnAuthorization", http.StatusUnauthorized,"error",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,res)
			return
		}

		claim,ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			res := helper.APIResponse("UnAuthorization", http.StatusUnauthorized,"error",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,res)
			return
		}

		userId := int(claim["user_id"].(float64))

		user, err := userService.GetUserById(userId)
		if err != nil {
			res := helper.APIResponse("UnAuthorization", http.StatusUnauthorized,"error",nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,res)
			return
		}

		c.Set("currentUser", user)
	}
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