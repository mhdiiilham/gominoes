package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/mhdiiilham/gominoes/api/app"
	db "github.com/mhdiiilham/gominoes/db"
	"github.com/mhdiiilham/gominoes/entity/user"
	"github.com/mhdiiilham/gominoes/pkg/jwt"
	"github.com/mhdiiilham/gominoes/pkg/password"
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
	passwordCrypter := password.NewBCryptHash()
	userManager := user.NewManager(userRepo, passwordCrypter)

	managers := &app.Managers{
		UserManager: userManager,
	}

	r := app.SetupApp(managers, jwtService, v, trans)
	r.Listen(":3000")
}
