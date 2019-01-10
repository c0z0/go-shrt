package main

import (
	"os"

	"github.com/c0z0/go-shrt/app"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := app.App{}
	app.Init()

	PORT := os.Getenv("PORT")

	if len(PORT) == 0 {
		PORT = "3000"
	}

	app.Run(":" + PORT)
}
