package main

import (
	"log"
	"net/http"
	"strings"
	"test_aac/auth"
	"test_aac/handler"
	"test_aac/helper"
	"test_aac/item"
	"test_aac/transaction"
	"test_aac/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Rcs150998:)@tcp(127.0.0.1:3306)/test_aac?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// test, err := auth.NewService().ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.WIiOEh6nkFOOumyKVfCJxCGIG_aGShfxegf_qwKcz1k")
	// if test.Valid {
	// 	fmt.Println("Valid")
	// } else {
	// 	fmt.Println("error")
	// }

	UserRepository := user.NewRepository(db)
	ItemRepository := item.NewReprository(db)
	TransRepository := transaction.NewRepository(db)

	UserService := user.NewService(UserRepository)
	AuthService := auth.NewService()
	ItemService := item.NewService(ItemRepository)
	TransService := transaction.NewService(TransRepository, ItemRepository)

	userHandler := handler.NewUserHandler(UserService, AuthService)
	itemHandler := handler.NewItemHandler(ItemService)
	transHandler := handler.NewTransHandler(TransService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images/", "./images")
	api := router.Group("/api/v1")

	api.POST("/sessions", userHandler.Login)
	api.GET("/users/fetch", authMiddleware(AuthService, UserService), userHandler.FetchUser)

	api.GET("/items", itemHandler.GetItem)
	api.GET("/items/:id", itemHandler.GetDetail)
	api.DELETE("/items/:id", authMiddleware(AuthService, UserService), itemHandler.DeleteItem)
	api.POST("/items", authMiddleware(AuthService, UserService), itemHandler.CreateItem)
	api.PUT("/items/:id", authMiddleware(AuthService, UserService), itemHandler.UpdateItem)

	api.GET("/transaction", authMiddleware(AuthService, UserService), transHandler.GetAllTransaction)
	api.POST("/transaction", authMiddleware(AuthService, UserService), transHandler.TransactionCreate)
	api.GET("/transaction/detail", authMiddleware(AuthService, UserService), transHandler.GetTransactionbyNumber)
	api.GET("/transaction/date", authMiddleware(AuthService, UserService), transHandler.GetTransactionbyDate)
	api.DELETE("/transaction/:id", authMiddleware(AuthService, UserService), transHandler.DeleteTrans)
	router.Run()
}

func authMiddleware(AuthService auth.Service, UserService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.JsonResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		//Bareer Token Header
		tokenString := ""
		codeToken := strings.Split(authHeader, " ")
		if len(codeToken) == 2 {
			tokenString = codeToken[1]
		}

		token, err := AuthService.ValidateToken(tokenString)
		if err != nil {
			response := helper.JsonResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.JsonResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := UserService.GetUserByID(userID)
		if err != nil {
			response := helper.JsonResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("codeUser", user)

	}
}
