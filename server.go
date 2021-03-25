package main

import (
	"fmt"
	"os"
	"shukai-api/infrastructure"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("env load error")
	}

	dsn := os.Getenv("DATABASE_URL")
	sqlHandler := infrastructure.NewSqlHandler(dsn)
	port := os.Getenv("PORT")
	router := infrastructure.NewRouting(sqlHandler, port)
	router.Run()
}
