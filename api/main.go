package main

import (
	"os"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/joho/godotenv"
	"github.com/mhdiiilham/gominoes/api/app"
	"github.com/mhdiiilham/gominoes/api/controllers"
	db "github.com/mhdiiilham/gominoes/db"
	"github.com/mhdiiilham/gominoes/entity/user"
	jwt "github.com/mhdiiilham/gominoes/pkg/jwt"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		panic("Translator not found!")
	}
	v := validator.New()
	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		panic(err)
	}
	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	_ = v.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} must be a valid email", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

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
