package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/mhdiiilham/gominoes/api/app"
	"github.com/mhdiiilham/gominoes/api/controllers"
	db "github.com/mhdiiilham/gominoes/db"
	"github.com/mhdiiilham/gominoes/entity/user"
	jwt "github.com/mhdiiilham/gominoes/pkg/jwt"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	v, trans := app.SetupValidator()

	client, err := db.NewMongoDBConnection(os.Getenv("MONGO_DB_USER"), os.Getenv("MONGO_DB_PASS"), os.Getenv("MONGO_DB"))
	if err != nil {
		panic(err)
	}
	jwtService := jwt.NewJWTService(os.Getenv("JWT_SECRET"), os.Getenv("APP_NAME"))
	userCollection := client.Database(os.Getenv("MONGO_DB")).Collection("users")
	userRepo := user.NewMongoDBRepository(userCollection)
	userManager := user.NewManager(userRepo)

	r := app.SetupApp()
	api := r.Group("/api")
	controllers.NewUserController(api, userManager, jwtService, v, trans)
	r.Listen(":3000")
}
