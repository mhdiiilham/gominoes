package main

import (
	"github.com/joho/godotenv"
	"github.com/mhdiiilham/gominoes/api/app"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	v, trans := app.SetupValidator()
	r := app.SetupApp(v, trans)
	r.Listen(":3000")
}
