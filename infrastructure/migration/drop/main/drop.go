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

	sqlHandler.Conn.Migrator().DropTable("notification_models")
	sqlHandler.Conn.Migrator().DropTable("approval_wait_models")
	sqlHandler.Conn.Migrator().DropTable("recruitment_users_models")
	sqlHandler.Conn.Migrator().DropTable("recruitment_conditions_models")
	sqlHandler.Conn.Migrator().DropTable("recruitment_models")
	sqlHandler.Conn.Migrator().DropTable("user_skill_models")
	sqlHandler.Conn.Migrator().DropTable("user_models")
}
