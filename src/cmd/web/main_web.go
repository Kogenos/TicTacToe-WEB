package main

import (
	"TicTacToe/internal/di"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/fx"
)

func main() {
	fx.New(di.App).Run()
}
